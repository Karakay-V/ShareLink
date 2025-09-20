<template>
  <div
    :class="`dropzone-wrapper ${isDragging ? 'dragging' : ''}`"
    :style="{ width: width, height: height }"
    @dragover="dragover"
    @dragleave="dragleave"
    @drop="drop"
  >
    <input
      type="file"
      name="file"
      id="fileInput"
      class="hidden-input"
      @change="onChange"
      ref="file"
      accept=".pdf,.jpg,.jpeg,.png"
    />

    <div class="info-overlay">
        <Button v-if="files.length === 0" label="Add File" width="150px" height="35px" />

        <label v-if="files.length === 0" for="fileInput" class="file-label">
            <div v-if="isDragging">Release to drop file here</div>
            <div v-else>Drag and drop your file here</div>
        </label>

        <!-- file list -->
        <ul v-else class="file-list">
            <li v-for="(file, index) in files" :key="index">
                <div    class="file-wrapper"
                        @click="removeFile(index)"
                >
                    <img class="file-icon" :src="FilePic" :alt="file.type" />
                    <span class="file-name">{{ file.name }}</span>
                </div>
            </li>
        </ul>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Button from "./Button.vue";
import FilePic from '../../assets/icons/file.icon.svg';
import ClosePic from '../../assets/icons/close.icon.svg';

export default defineComponent({
    name: "DragAndDropArea",
    components: { Button },
    data() {
        return {
            isDragging: false,
            files: [] as File[],
            FilePic,
            ClosePic,
        };
    },
    props: {
        width: {
            type: String,
            required: false,
            default: "100%",
        },
        height: {
            type: String,
            required: false,
            default: "180px",
        },
    },
    methods: {
        onChange() {
            const input = this.$refs.file as HTMLInputElement;
            if (input && input.files && input.files[0]) {
            this.files = [input.files[0]]; // replace instead of push (single file)
            }
        },
        dragover(e: DragEvent) {
            e.preventDefault();
            this.isDragging = true;
        },
        dragleave() {
            this.isDragging = false;
        },
        drop(e: DragEvent) {
            e.preventDefault();
            if (e.dataTransfer?.files && e.dataTransfer.files[0]) {
                this.files = [e.dataTransfer.files[0]]; // replace
            }
            this.isDragging = false;
        },
        removeFile(index: number) {
            this.files.splice(index, 1);
            const input = this.$refs.file as HTMLInputElement;
            if (input) input.value = ""; // clear input value
        },
    },
});
</script>

<style lang="scss" scoped>
@use "/src/assets/styles/colors" as *;
@use "/src/assets/styles/image";
@use "/src/assets/styles/fonts";
@use "/src/assets/styles/shadows";

.dropzone-wrapper {
  border-radius: 30px;
  border: 2px dashed $border-gray;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: 0.3s ease background-color, 0.3s ease box-shadow,
    0.3s ease color, 0.3s ease border-color;

  &.dragging {
    background-color: $background-gray;
    border-color: $border-orange;

    .file-label {
      color: $text-orange;
    }
  }

  input {
    opacity: 0;
    width: 100%;
    height: 100%;
    position: absolute;
    top: 0;
    left: 0;
    cursor: pointer;
  }

  .info-overlay {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 15px;

    .file-label {
      pointer-events: none;
      @include fonts.noto-font(600);
      @include fonts.responsive-font(12, 12, 1440);
      @include fonts.prevent-selecting;
      transition: 0.3s ease color;
      text-align: center;
      color: $text-gray;
    }

    .file-list {
      list-style: none;
      padding: 0;
      margin: 0;
      width: 100%;
      max-width: 300px;

      li {
         .file-wrapper {
            border: 2px solid $border-gray;
            background: $background-white;
            border-radius: 10px;
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: space-between;
            padding: 4px 6px;
            z-index: 2;

            .file-icon {
                width: 70px;
            }

            .file-name {
                @include fonts.noto-font(600);
                @include fonts.responsive-font(13, 13, 1440);
                @include fonts.prevent-selecting;
                color: $text-gray;
                text-align: center;
            }
        }
      }
    }
  }
}
</style>
