<template>
  <div class="textlabel">
    {{ $t("repository.choose-git-provider-description") }}
  </div>
  <div class="mt-4 flex flex-wrap space-x-2 gap-y-2">
    <template v-for="(vcs, index) in vcsList" :key="index">
      <NButton type="default" @click.prevent="selectVCS(vcs)">
        <template #icon>
          <VCSIcon custom-class="h-6" :type="vcs.type" />
        </template>
        <span>{{ vcs.title }}</span>
      </NButton>
    </template>
  </div>
  <div class="mt-4 textinfolabel">
    <template v-if="canManageVCSProvider">
      <i18n-t keypath="repository.choose-git-provider-visit-workspace">
        <template #workspace>
          <router-link
            class="normal-link"
            :to="{ name: WORKSPACE_ROUTE_GITOPS }"
          >
            {{ $t("common.workspace") }} - {{ $t("common.gitops") }}
          </router-link>
        </template>
      </i18n-t>
    </template>
    <template v-else>
      {{ $t("repository.choose-git-provider-contact-workspace-admin") }}
    </template>
  </div>
</template>

<script lang="ts">
export default { name: "RepositoryVCSProviderPanel" };
</script>

<script setup lang="ts">
import { reactive, computed, watchEffect } from "vue";
import { hasWorkspacePermissionV2 } from "@/utils";
import { useVCSProviderStore } from "@/store";
import type { VCSProvider } from "@/types/proto/v1/vcs_provider_service";
import { WORKSPACE_ROUTE_GITOPS } from "@/router/dashboard/workspaceRoutes";
import { NButton } from "naive-ui";
import { VCSIcon } from "./VCS";

interface LocalState {
  selectedVCS?: VCSProvider;
}

const emit = defineEmits<{
  (event: "next"): void;
  (event: "set-vcs", payload: VCSProvider): void;
}>();

const vcsV1Store = useVCSProviderStore();
const state = reactive<LocalState>({});

const prepareVCSList = () => {
  vcsV1Store.getOrFetchVCSList();
};

watchEffect(prepareVCSList);

const vcsList = computed(() => {
  return vcsV1Store.vcsList;
});

const canManageVCSProvider = computed(() => {
  return hasWorkspacePermissionV2("bb.vcsProviders.list");
});

const selectVCS = (vcs: VCSProvider) => {
  state.selectedVCS = vcs;
  emit("set-vcs", vcs);
  emit("next");
};
</script>
