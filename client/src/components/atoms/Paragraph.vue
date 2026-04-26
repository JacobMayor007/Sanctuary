<script setup lang="ts">
import clsx from 'clsx';
import { twMerge } from 'tailwind-merge';

const props = defineProps<{
  title: string;
  paraClass?: string;
}>();

const paraClasses = twMerge(clsx('font-manrope text-[#47645C]', props.paraClass));
</script>

<template>
  <p :class="paraClasses">
    {{ title }}
  </p>
</template>
