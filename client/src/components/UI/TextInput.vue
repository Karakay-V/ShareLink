<template>
  <div :class="`wrapper ${message ? '' : 'messages-hidden'}`">
    <div class="input_container">
      <input
        :type="type"
        :value="data"
        :class="`${isFocused ? 'focused' : ''} ${errorMessage !== '' ? 'error' : ''}`"
        :maxlength="maxLength"
        @input="handleInput"
        @focus="handleFocus"
        @blur="handleBlur"
      />

      <label :class="`floating_label ${ isFocused || data ? 'active' : '' }`">
        {{ placeholder }}
      </label>
    </div>

    <label :class="`message ${errorMessage !== '' ? 'active' : ''}`">{{ errorMessage }}</label>
  </div>
</template>

<script lang="ts">
  import { defineComponent } from 'vue'
  import { InputDataTypes } from '../../types/input-data-types';

  export default defineComponent({
    name: "TextInput",
    data() {
      return {
        InputDataTypes,
        isFocused: false,
        errorMessage: "",
      };
    },
    props: {
      data: {
        type: String,
        required: true,
      },
      type: {
        type: String,
        default: InputDataTypes.Text,
        required: false,
        validator: (value: string) =>
            (Object.values(InputDataTypes) as string[]).includes(value),
      },
      required: {
        type: Boolean,
        default: false,
        required: false,
      },
      placeholder: {
        type: String,
        default: "",
        required: false,
      },
      maxLength: {
        type: Number,
        default: 50,
        required: false,
      },
      message: {
        type: Boolean,
        required: false,
        default: true,
      },
    },
    emits: ['update:data', 'focus-change'],
    methods: {
      handleInput(event: Event) {
        const target = event.target as HTMLInputElement | null;
        if (target) {
          const value = target.value;
          this.$emit('update:data', value);
        }
      },
      handleFocus(event: Event) {
        this.isFocused = true;
        this.$emit('focus-change', true);
      },
      handleBlur(event: Event) {
        this.isFocused = false;
        const target = event.target as HTMLInputElement;
        this.validateInput(target.value);
      },
      validateInput(value: string) {
        if (this.required || this.data !== "") {
        /***      If field is required or user typed some characters into input      ***/
          switch (this.type) {
            case "text":
              if (!value.trim() && this.required == true) {
                this.errorMessage = "This field cannot be empty";
                return false;
              }
              break;
            case "email":
              if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value)) {
                this.errorMessage = "Please enter a valid email address";
                return false;
              }
              break;
            case "tel":
              if (!/^\+?[0-9]{10,15}$/.test(value)) {
                this.errorMessage = "Please enter a valid phone number";
                return false;
              }
              break;
            case "url":
              try {
                new URL(value);
              } catch (_) {
                this.errorMessage = "Please enter a valid URL";
                return false;
              }
              break;
          }
        }

        /***      If all is good:      ***/
        this.errorMessage = "";
        return true;
      }
    }
  })
</script>

<style scoped lang="scss">
  @use "/src/assets/styles/fonts" as *;
  @use "/src/assets/styles/colors" as *;
  @use "/src/assets/styles/image";
  @use "/src/assets/styles/shadows";

  .wrapper {
    display: flex;
    flex-direction: column;
    gap: 12px;
    padding: 5px;
    width: calc(100% - 5px - 5px);
    max-width: 500px;

    &.messages-hidden {
        gap: 0;
        
        .message {
            display: none;
        }
    }
  }

  .input_container {
    position: relative;
    display: inline-block;
    cursor: text;
    width: 100%;

    input {
      // Minus 44px for paddings and 4px for borders
      width: calc(100% - 44px - 4px);
      padding: 12px 15px;
      border-radius: 50px;
      border: 2px solid transparent;
      @include noto-font(400);
      @include shadows.box-shadow-medium;
      @include responsive-font(13, 13, 1440);
      transition: 0.3s ease background-color, 0.3s ease box-shadow, 0.3s ease color, 0.3s ease border-color;
      background-color: $background-white;
      color: $text-black;

      &.focused {
        @include shadows.box-shadow-double;
      }

      &.error {
        color: $text-error;
        border-color: $text-error;
      }

      &:focus,
      &:focus-visible,
      &:focus-within {
        outline: none;
      }
    }

    .floating_label {
      position: absolute;
      left: 15px;
      top: 12px;
      color: $text-placeholder-gray;
      opacity: 1;
      transition: 0.3s ease opacity, 0.3s ease color;
      pointer-events: none;

      @include noto-font(400);
      @include responsive-font(16, 16, 1440);
      @include hide-overflowed-text;
      width: calc(100% - 48px);


      &.active {
        color: $text-black;
        opacity: 0;
      }

      &.disabled {
        color: $text-placeholder-gray !important;
      }
    }

    &:hover {
      input {
        background-color: $background-white;
      }
    }

  }

  .message {
    @include noto-font(400);
    @include responsive-font(16, 16, 1440);
    color: $text-error;
    transition: 0.3s ease opacity;
    opacity: 0;

    &.active {
        opacity: 1;
    }
  }
</style>
