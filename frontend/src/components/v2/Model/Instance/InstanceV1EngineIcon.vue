<template>
  <NTooltip :disabled="!tooltip || !instance.engineVersion">
    <template #trigger>
      <div class="relative w-4 shrink-0" v-bind="$attrs">
        <EngineIcon :engine="instance.engine" />
        <div
          v-if="showStatus"
          class="bg-green-400 border-surface-high rounded-full absolute border-2"
          style="bottom: -3px; height: 9px; right: -3px; width: 9px"
        />
      </div>
    </template>
    <span>{{ instance.engineVersion }}</span>
  </NTooltip>
</template>

<script lang="ts" setup>
import { NTooltip } from "naive-ui";
import type { PropType } from "vue";
import { EngineIcon } from "@/components/Icon";
import type {
  Instance,
  InstanceResource,
} from "@/types/proto/v1/instance_service";

defineProps({
  instance: {
    required: true,
    type: Object as PropType<Instance | InstanceResource>,
  },
  showStatus: {
    type: Boolean,
    default: false,
  },
  tooltip: {
    type: Boolean,
    default: true,
  },
});
</script>
