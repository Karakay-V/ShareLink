<template>
  <!-- :style="{ width: width, height: height }" -->
  <div
    :class="`dropzone-wrapper ${isDragging ? 'dragging' : ''}`"
    @dragover="dragover"
    @dragleave="dragleave"
    @drop="drop"
  >
    <input
      type="file"
      name="file"
      multiple
      id="fileInput"
      class="hidden-input"
      @change="onChange"
      ref="file"
      accept=".pdf,.jpg,.jpeg,.png"
    />

    <div class="info-overlay">
        <Button v-if="files.length === 0" 
                label="Add File" 
                width="150px" 
                height="35px"
                @click="openFileDialog" />

        <label v-if="files.length === 0" for="fileInput" class="file-label">
            <div v-if="isDragging">Release to drop file here</div>
            <div v-else>Drag and drop your file here</div>
        </label>

        <!-- file list -->
        <ul v-else class="file-list">
            <li v-for="(file, index) in files" :key="index">
                <div    class="file-wrapper"
                        @click="removeFile(index)">
                    <div class="remove-wrapper">
                      <img class="remove-icon" :src="ClosePic" alt="Remove" />
                    </div>
                    
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
            default: "max-content",
        },
    },
    methods: {
        onChange() {
           const input = this.$refs.file as HTMLInputElement;
          if (input && input.files) {
              const selected = Array.from(input.files);
              this.files = [...this.files, ...selected].slice(0, 5);
              input.value = "";
          }
        },
        dragover(_event: DragEvent) {
            _event.preventDefault();
            this.isDragging = true;
        },
        dragleave() {
            this.isDragging = false;
        },
        drop(_event: DragEvent) {
            _event.preventDefault();
            if (_event.dataTransfer?.files) {
                const dropped = Array.from(_event.dataTransfer.files);
                this.files = [...this.files, ...dropped].slice(0, 5);
            }
            this.isDragging = false;
        },
        removeFile(index: number) {
            this.files.splice(index, 1);
            const input = this.$refs.file as HTMLInputElement;
            if (input) input.value = ""; // clear input value
        },
        openFileDialog() {
            const input = this.$refs.file as HTMLInputElement;
            input?.click();
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
  min-height: 140px;
  width: calc(100% - 20px - 20px);
  padding: 20px 20px;
  transition: 0.3s ease background-color, 0.3s ease box-shadow,
    0.3s ease color, 0.3s ease border-color, 0.3s ease transform;

  &.dragging {
    background-color: $background-gray;
    border-color: $border-orange;
    transform: scale(1.025);

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
    z-index: 2;

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
      display: flex;
      align-items: start;
      justify-content: space-between;
      gap: 10px 20px;
      flex-wrap: wrap;

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
            cursor: pointer;
            position: relative;
            transition: ease 0.3s all;
            width: 110px;
            height: 100px;

            &:hover {
                background-color: $background-gray;
                border-color: $border-danger;
                transform: scale(1.025);

                .remove-wrapper {
                    opacity: 1;
                }
            }

            .remove-wrapper {
                transition: ease 0.3s opacity;
                opacity: 0;
                width: 100%;
                height: 100%;
                position: absolute;
                top: 0;
                left: 0;
                display: flex;
                justify-content: center;
                align-items: center;
                
                .remove-icon {
                    width: 70%;
                    height: 70%;
                }
            }

            .file-icon {
                width: 70px;
            }

            .file-name {
                @include fonts.noto-font(600);
                @include fonts.responsive-font(13, 13, 1440);
                @include fonts.prevent-selecting;
                @include fonts.hide-overflowed-text;
                color: $text-gray;
                text-align: center;
            }
        }
      }
    }
  }
}
</style>
