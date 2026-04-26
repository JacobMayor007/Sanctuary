<script setup lang="ts">
import Button from '../atoms/Button.vue';
import InputGroup from '../molecules/InputGroup.vue';
</script>

<template>
  <form action="" class="flex flex-col gap-5 2xl:gap-5">
    <InputGroup
      class="mb-5"
      parentClass="gap-4"
      label="full name"
      type="text"
      placeholder="Eleanor Thorne"
    />
    <InputGroup
      class="mb-5"
      parentClass="gap-4"
      label="Email Address"
      type="email"
      placeholder="eleanorthorne@gmail.com"
    />
    <InputGroup
      class="mb-5"
      parentClass="gap-4"
      label="password"
      type="password"
      placeholder="***********"
      inputClass="tracking-[0.2em]"
    />
    <Button
      title="Begin My Ritual"
      buttonClass="py-7 rounded-xl bg-[#5F7D74] text-white text-[16px]"
    />
  </form>
</template>
