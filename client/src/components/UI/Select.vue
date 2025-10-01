<template>
  <div
    class="ui-select"
    :class="{ open: isOpen }"
    @keydown.stop.prevent="onKeydown"
    tabindex="0"
    ref="root"
    role="combobox"
    :aria-expanded="isOpen"
  >
    <button
      type="button"
      class="select-toggle"
      @click="toggle"
      :aria-haspopup="true"
      :aria-controls="listId"
    >
      <span class="label">
        <span v-if="selectedLabel">{{ selectedLabel }}</span>
        <span v-else class="placeholder">{{ placeholder }}</span>
      </span>

      <svg class="chev" viewBox="0 0 24 24" aria-hidden="true">
        <path d="M7 10l5 5 5-5z" fill="currentColor"/>
      </svg>
    </button>

    <ul
      v-if="isOpen"
      :id="listId"
      class="options"
      role="listbox"
      :aria-activedescendant="activeId"
      @mousedown.prevent
    >
      <li
        v-for="(it, idx) in itemsLocal"
        :key="itemKey(it, idx)"
        :id="optionId(idx)"
        class="option"
        :class="{ highlighted: idx === highlightedIndex }"
        role="option"
        :aria-selected="(isSelected(it))"
        @click="selectItem(it)"
        @mouseenter="highlightedIndex = idx"
      >
        <slot name="option" :item="it">
          {{ itemLabel(it) }}
        </slot>
      </li>

      <li v-if="itemsLocal.length === 0" class="option empty">
        No options
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, watch, onMounted, onBeforeUnmount, computed } from "vue";

export default defineComponent({
  name: "Select",
  props: {
    // options array (array of strings or objects)
    items: {
      type: Array as () => any[],
      default: () => []
    },
    // v-model for selected item
    modelValue: {
      required: false
    },
    // placeholder text
    placeholder: {
      type: String,
      default: "Select..."
    },
    // key/label accessors (when items are objects)
    labelKey: {
      type: String,
      default: "name"
    },
    valueKey: {
      type: String,
      default: "id"
    },
    // use custom id prefix for DOM ids
    idPrefix: {
      type: String,
      default: "ui-select"
    }
  },
  emits: ["update:modelValue", "update:items", "change"],
  setup(props, { emit }) {
    const isOpen = ref(false);
    const highlightedIndex = ref(-1);
    const root = ref<HTMLElement | null>(null);

    // local copy so parent can v-model:items (update:items)
    const itemsLocal = ref<any[]>([...props.items]);

    watch(() => props.items, (v) => {
      itemsLocal.value = [...v || []];
      if (highlightedIndex.value >= itemsLocal.value.length) highlightedIndex.value = -1;
    }, { deep: true });

    // when parent uses v-model:items -> update itemsLocal via emitted update:items
    // provide method to update items from inside if needed
    function setItems(next: any[]) {
      itemsLocal.value = [...next];
      emit("update:items", itemsLocal.value);
    }

    // selected item (modelValue)
    const selected = ref<any>(props.modelValue ?? null);
    watch(() => props.modelValue, (v) => { selected.value = v; });

    function updateSelected(v: any) {
      selected.value = v;
      emit("update:modelValue", v);
      emit("change", v);
    }

    // textual label shown
    function itemLabel(it: any) {
      if (it == null) return "";
      if (typeof it === "string" || typeof it === "number") return String(it);
      if (props.labelKey && typeof it === "object" && it[props.labelKey] !== undefined) {
        return String(it[props.labelKey]);
      }
      return String(it);
    }

    const selectedLabel = computed(() => {
      return selected.value != null ? itemLabel(selected.value) : "";
    });

    function isSelected(it: any) {
      if (selected.value == null) return false;
      if (typeof selected.value === "object" && typeof it === "object" && props.valueKey) {
        return selected.value[props.valueKey] === it[props.valueKey];
      }
      return selected.value === it;
    }

    function itemKey(it: any, idx: number) {
      if (it == null) return `item-${idx}`;
      if (typeof it === "object" && props.valueKey && it[props.valueKey] !== undefined) {
        return String(it[props.valueKey]);
      }
      return `${String(it)}-${idx}`;
    }

    function optionId(idx: number) {
      return `${props.idPrefix}-option-${idx}`;
    }

    const listId = `${props.idPrefix}-list`;
    const activeId = computed(() => (highlightedIndex.value >= 0 ? optionId(highlightedIndex.value) : ""));

    function open() {
      isOpen.value = true;
      // highlight current selected index
      const idx = itemsLocal.value.findIndex((it) => isSelected(it));
      highlightedIndex.value = idx >= 0 ? idx : 0;
    }
    function close() {
      isOpen.value = false;
      highlightedIndex.value = -1;
    }
    function toggle() {
      if (isOpen.value) close(); else open();
    }

    function selectItem(it: any) {
      updateSelected(it);
      close();
    }

    // keyboard navigation
    function onKeydown(e: KeyboardEvent) {
      if (e.key === "ArrowDown") {
        if (!isOpen.value) { open(); return; }
        highlightedIndex.value = Math.min(itemsLocal.value.length - 1, Math.max(0, highlightedIndex.value + 1));
      } else if (e.key === "ArrowUp") {
        if (!isOpen.value) { open(); return; }
        highlightedIndex.value = Math.max(0, highlightedIndex.value - 1);
      } else if (e.key === "Enter") {
        if (isOpen.value && highlightedIndex.value >= 0 && highlightedIndex.value < itemsLocal.value.length) {
          selectItem(itemsLocal.value[highlightedIndex.value]);
        } else {
          open();
        }
      } else if (e.key === "Escape") {
        close();
      }
    }

    // click outside to close
    function onDocumentClick(e: MouseEvent) {
      if (!root.value) return;
      if (!root.value.contains(e.target as Node)) close();
    }
    onMounted(() => document.addEventListener("click", onDocumentClick));
    onBeforeUnmount(() => document.removeEventListener("click", onDocumentClick));

    // expose some API if parent uses ref
    return {
      isOpen,
      highlightedIndex,
      itemsLocal,
      root,
      listId,
      activeId,
      toggle,
      open,
      close,
      selectItem,
      itemLabel,
      itemKey,
      optionId,
      onKeydown,
      selectedLabel,
      setItems,
      updateSelected,
      isSelected
    };
  }
});
</script>

<style lang="scss" scoped>
@use "/src/assets/styles/colors" as *;
@use "/src/assets/styles/fonts";

.ui-select {
  width: 100%;
  border-radius: 30px;
  background: $background-white;
  border: 2px solid $border-gray;
  position: static;
  transition: 0.2s ease all;
  outline: none;

  &:focus-within {
    box-shadow: 0 6px 18px rgba(0,0,0,0.06);
    transform: translateZ(0);
  }

  .select-toggle {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 12px;
    padding: 12px 18px;
    width: 100%;
    background: transparent;
    border: 0;
    cursor: pointer;
    font-family: inherit;
    color: $text-black;
    border-radius: 28px;
    text-align: left;

    .label {
      display: inline-block;
      @include fonts.noto-font(500);
      @include fonts.responsive-font(16, 16, 1440);
      color: $text-black;
    }

    .placeholder {
      color: $text-placeholder-gray;
    }

    .chev {
      width: 20px;
      height: 20px;
      color: $border-gray;
      transition: 0.18s transform;
    }
  }

  &.open {
    border-color: $border-orange;

    .chev {
      transform: rotate(180deg);
      color: $border-orange;
    }
  }

  .options {
    position: absolute;
    top: calc(100%);
    left: 0;
    width: calc(100% - 8px - 8px);
    max-height: 320px;
    overflow: auto;
    background: $background-white;
    border: 1px solid $border-gray;
    border-radius: 10px;
    padding: 8px;
    box-shadow: 0 8px 30px rgba(0,0,0,0.08);
    z-index: 4;

    .option {
      padding: 10px 12px;
      border-radius: 8px;
      cursor: pointer;
      display: flex;
      align-items: center;
      gap: 8px;
      color: $text-black;
      @include fonts.noto-font(500);
      @include fonts.responsive-font(16, 16, 1440);

      &.highlighted {
        background: $background-gray;
        color: $text-orange;
      }

      &.empty {
        color: $text-gray;
        text-align: center;
      }
    }
  }
}
</style>
