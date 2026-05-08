<script setup lang="ts">
import InputGroup from '../molecules/InputGroup.vue';
import IconEyeOpen from '../../../public/icons/IconEyeOpen.vue';
import IconEyeClose from '../../../public/icons/IconEyeClose.vue';
import Button from '../atoms/Button.vue';
import Heading from '../atoms/Heading.vue';
import { ref } from 'vue';

const seePassword = ref(false);
</script>

<template>
  <form @submit.prevent="" class="flex flex-col gap-10">
    <InputGroup type="email" id="email" label="Email" placeholder="ritual@sanctuary.com" />
    <InputGroup :type="seePassword ? 'text' : 'password'" placeholder="* * * * * * * *">
      <template #right-icon>
        <IconEyeOpen
          v-if="seePassword"
          class="cursor-pointer active:scale-90"
          @click="seePassword = false"
        />
        <IconEyeClose v-else class="cursor-pointer active:scale-90" @click="seePassword = true" />
      </template>
    </InputGroup>
    <Heading
      @click="$router.push('/forgot-password')"
      title="Forgot Password?"
      class="font-manrope cursor-pointer text-right text-[11px] font-bold tracking-widest text-[#47645C] uppercase"
    />

    <Button title="Sign In" type="submit" />
  </form>
</template>
