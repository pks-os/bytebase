import { TemplateType } from "@/plugins";
import {
  DatabaseId,
  IssueCreate,
  MigrationContext,
  MigrationDetail,
  MigrationType,
  RollbackDetail,
  UNKNOWN_ID,
} from "@/types";
import {
  findDatabaseListByQuery,
  BuildNewIssueContext,
  ESTABLISH_BASELINE_SQL,
  VALIDATE_ONLY_SQL,
} from "../common";
import { IssueCreateHelper } from "./helper";

export const buildNewStandardIssue = async (
  context: BuildNewIssueContext
): Promise<IssueCreate> => {
  const helper = new IssueCreateHelper(context);

  await helper.prepare();

  // standard single-stage or multi-stage issue is generated by specifying
  // databaseId for each stage
  const databaseList = findDatabaseListByQuery(context);
  const templateType = context.route.query.template as TemplateType;
  let migrationType: MigrationType = "MIGRATE";
  if (templateType === "bb.issue.database.data.update") {
    migrationType = "DATA";
  }
  if (templateType === "bb.issue.database.schema.baseline") {
    migrationType = "BASELINE";
  }
  const statement =
    migrationType === "BASELINE" ? ESTABLISH_BASELINE_SQL : VALIDATE_ONLY_SQL;
  const createContext: MigrationContext = {
    detailList: databaseList.map((db, index) => {
      const detail: MigrationDetail = {
        migrationType: migrationType,
        databaseId: Number(db.uid),
        statement,
        sheetId: UNKNOWN_ID,
        earliestAllowedTs: 0,
      };
      return detail;
    }),
  };
  maybeApplyRollbackParams(createContext, context.route);

  helper.issueCreate!.createContext = createContext;
  await helper.validate();

  // clean up createContext for standard issues
  helper.issueCreate!.createContext = {};

  return helper.generate();
};

// Try to find out the relationship between databaseId and rollback issue/task
// Id from the URL query.
export const getRollbackTaskMappingFromQuery = (
  route: BuildNewIssueContext["route"]
) => {
  const mapping = new Map<DatabaseId, RollbackDetail>();

  const { query } = route;
  const databaseIdListInQuery = (query.databaseList as string) || "";
  const databaseIdList = databaseIdListInQuery
    .split(",")
    .map<DatabaseId>((maybeId) => parseInt(maybeId, 10));

  const rollbackIssueIdInQuery = (query.rollbackIssueId as string) || "";
  const issueId = parseInt(rollbackIssueIdInQuery, 10) || UNKNOWN_ID;
  if (issueId === UNKNOWN_ID) {
    return mapping;
  }

  const rollbackTaskIdListInQuery = (query.rollbackTaskIdList as string) || "";
  const taskIdList = rollbackTaskIdListInQuery
    .split(",")
    .map((maybeId) => parseInt(maybeId, 10) || UNKNOWN_ID);

  databaseIdList.forEach((databaseId, index) => {
    const taskId = taskIdList[index] || UNKNOWN_ID;
    if (taskId !== UNKNOWN_ID) {
      mapping.set(databaseId, {
        issueId,
        taskId,
      });
    }
  });
  return mapping;
};

export const maybeApplyRollbackParams = (
  migrationContext: MigrationContext,
  route: BuildNewIssueContext["route"]
) => {
  const mapping = getRollbackTaskMappingFromQuery(route);
  migrationContext.detailList.forEach((detail) => {
    const { databaseId } = detail;
    if (!databaseId || databaseId === UNKNOWN_ID) return;

    const rollbackDetail = mapping.get(databaseId);
    if (!rollbackDetail) return;

    detail.rollbackDetail = rollbackDetail;
  });
};