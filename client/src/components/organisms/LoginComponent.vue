<script setup lang="ts">
import InputGroup from '../molecules/InputGroup.vue';
import IconEyeOpen from '../../../public/icons/IconEyeOpen.vue'
</script>

<template>
  <form @submit.prevent="" class="flex flex-col gap-10">
    <InputGroup type="email" id="email" label="Email" placeholder="ritual@sanctuary.com" />
    <InputGroup type="password" placeholder="* * * * * * * *">
      <template #right-icon>
        <IconEyeOpen class="cursor-pointer active:scale-90 " />
      </template>
    </InputGroup>
    <h3
      class="font-manrope text-right text-[11px] font-bold tracking-widest text-[#47645C] uppercase"
    >
      Forgot Password?
    </h3>
    <button
      class="font-manrope rounded-sm border-[1px] py-4.5 text-sm font-bold tracking-widest text-[#47645C] uppercase"
    >
      Sign In
    </button>
  </form>
</template>
