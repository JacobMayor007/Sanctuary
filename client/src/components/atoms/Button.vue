<script setup lang="ts">
import clsx from 'clsx';
import { twMerge } from 'tailwind-merge';

const props = defineProps<{
  title: string;
  buttonClass?: string;
}>();

const buttonClasses = twMerge(
  clsx(
    'font-manrope rounded-sm border-[1px] py-4.5 text-sm font-bold tracking-widest text-[#47645C] uppercase',
    props.buttonClass,
  ),
);
</script>

<template>
  <button :class="buttonClasses">
    {{ title }}
  </button>
</template>
