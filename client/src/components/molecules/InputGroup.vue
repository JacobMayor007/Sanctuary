<script setup lang="ts">
import { twMerge } from 'tailwind-merge';
import { clsx } from 'clsx';
import { useSlots } from 'vue';
const slots = useSlots();
const props = defineProps<{
  type: string;
  label?: string;
  placeholder?: string;
  id?: string;
  labelClass?: string;
  inputClass?: string;
  parentClass?: string;
}>();

const labelClasses = twMerge(
  clsx(
    'font-manrope text-[10px] font-bold tracking-[0.25em] text-[#47645C]/60 uppercase',
    props.labelClass,
  ),
);

const inputClasses = twMerge(
  clsx(
    'font-manrope placeholder:font-manrope h-11 w-full border-b-[1px] border-[#C1C8C4] text-sm text-[#8e9792] outline-none placeholder:text-[#1C1C18]/20',
    slots['left-icon'] && 'pl-6.5',
    slots['right-icon'] && 'pr-6.5',
    props.inputClass,
  ),
);

const parentClasses = twMerge(clsx('flex flex-col gap-1'), props.parentClass);

const emit = defineEmits<{
  change: [value: string];
}>();
</script>

<template>
  <div :class="parentClasses">
    <label :for="id" :class="labelClasses">{{ label }}</label>
    <div class="relative flex items-center">
      <div v-if="$slots['left-icon']" class="absolute left-0 bg-white">
        <slot name="left-icon" />
      </div>

      <input
        :type="type"
        :id="id"
        :placeholder="placeholder"
        :class="inputClasses"
        required
        @change="emit('change', ($event.target as HTMLInputElement).value)"
      />

      <div v-if="$slots['right-icon']" class="absolute right-0">
        <slot name="right-icon" />
      </div>
    </div>
  </div>
</template>
