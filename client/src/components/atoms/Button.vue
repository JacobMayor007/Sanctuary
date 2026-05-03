<script setup lang="ts">
import clsx from 'clsx';
import { twMerge } from 'tailwind-merge';

const props = defineProps<{
  title: string;
  buttonClass?: string;
  type: 'submit' | 'button' | 'reset';
}>();

const buttonClasses = twMerge(
  clsx(
    'font-manrope cursor-pointer rounded-sm border-[1px] py-4.5 text-sm font-bold tracking-widest text-[#47645C] uppercase active:scale-95',
    props.buttonClass,
  ),
);
</script>

<template>
  <button :type="type ?? 'button'" :class="buttonClasses">{{ title }}</button>
</template>
