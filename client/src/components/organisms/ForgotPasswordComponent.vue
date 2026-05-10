<script setup lang="ts">
import Button from '@/components/atoms/Button.vue';
import InputGroup from '@/components/molecules/InputGroup.vue';
</script>

<template>
  <form action="" class="flex flex-col gap-10">
    <InputGroup type="email" label="Email Address" placeholder="name@example.com" />
    <Button title="Send Reset Link" type="submit" button-class="bg-[#5F7D74] text-white" />
  </form>
</template>
