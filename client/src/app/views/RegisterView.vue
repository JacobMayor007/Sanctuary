<script setup lang="ts">
import Button from '@/components/atoms/Button.vue';
import Heading from '@/components/atoms/Heading.vue';
import Paragraph from '@/components/atoms/Paragraph.vue';
import Title from '@/components/atoms/title.vue';
import InputGroup from '@/components/molecules/InputGroup.vue';
import RegisterComponent from '@/components/organisms/RegisterComponent.vue';
</script>

<template>
  <main class="flex h-screen grid-cols-2 md:grid">
    <div
      class="hidden flex-1 flex-col items-center justify-center gap-4 bg-gradient-to-r from-[#F6F3ED] to-[#FCF9F3] text-center md:flex md:px-20"
    >
      <h1 class="font-libertinus text-5xl tracking-wide">The Tactile Breath</h1>
      <p class="font-manrope text-[12px] tracking-[0.2em] text-[#414846] uppercase">
        A journey into stillness begins here.
      </p>
    </div>
    <div class="flex-1 flex-col bg-[#FCF9F3] px-8 py-12 md:flex">
      <div class="flex flex-1 flex-col justify-around lg:px-20">
        <Title />
        <div>
          <h1 class="font-libertinus mb-8 text-5xl">Create Account</h1>
          <p class="font-manrope tracking-widest text-[#414846]">
            Join our curated community of wellness and international living.
          </p>
        </div>
        <RegisterComponent />
        <div class="flex flex-row justify-between">
          <Heading title="Already have an account?" headingClass="text-lg" />
          <Paragraph title="Sign In " paraClass="font-bold cursor-pointer" />
        </div>
      </div>
    </div>
  </main>
</template>
