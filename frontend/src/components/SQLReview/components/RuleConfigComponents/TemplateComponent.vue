<template>
  <InputWithTemplate
    :template-list="templateList"
    :value="value"
    :disabled="disabled"
    @change="$emit('update:value', $event)"
  />
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { useI18n } from "vue-i18n";
import {
  InputWithTemplate,
  type Template,
} from "@/components/InputWithTemplate";
import type { RuleConfigComponent, TemplatePayload } from "@/types";
import { getRuleLocalizationKey } from "@/types";

const props = defineProps<{
  ruleType: string;
  config: RuleConfigComponent;
  value: string;
  disabled: boolean;
}>();

defineEmits<{
  (event: "update:value", value: string): void;
}>();

const { t } = useI18n();

const templateList = computed(() => {
  const { ruleType, config } = props;
  const payload = config.payload as TemplatePayload;
  return payload.templateList.map<Template>((id) => ({
    id,
    description: t(
      `sql-review.rule.${getRuleLocalizationKey(ruleType)}.component.${
        config.key
      }.template.${id}`
    ),
  }));
});
</script>
