<script setup lang="ts">
import clsx from 'clsx';
import { twMerge } from 'tailwind-merge';

const props = defineProps<{
  title: string;
  headingClass?: string;
}>();

const headingClasses = twMerge(clsx('font-libertinus'), props.headingClass);
</script>
<template>
  <h1 :class="headingClasses">
    {{ title }}
  </h1>
</template>
