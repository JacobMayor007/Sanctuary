<script setup lang="ts">
import Heading from '@/components/atoms/Heading.vue';
import Paragraph from '@/components/atoms/Paragraph.vue';
import TitleHeader from '@/components/atoms/TitleHeader.vue';
import IconArrow from '../../../public/icons/IconArrow.vue';
import ForgotPasswordComponent from '@/components/organisms/ForgotPasswordComponent.vue';
</script>

<template>
  <main class="flex h-screen flex-col items-center gap-32 p-8">
    <TitleHeader />
    <div class="flex flex-col gap-8">
      <Heading title="Reset Password" class="font-libertinus text-center text-5xl text-[#1C1C18]" />
      <Paragraph
        title="Enter your email address and we'll send you a link to
          reset your password."
        class="font-manrope text-center tracking-wide"
      />
      <ForgotPasswordComponent />
      <div
        @click="$router.push('/sign-in')"
        class="flex flex-row items-center justify-center gap-8"
      >
        <IconArrow width="18" height="18" />
        <Paragraph title="Back to Sign In" para-class="text-center cursor-pointer" />
      </div>
    </div>
  </main>
</template>
