<script setup lang="ts">
import Heading from '@/components/atoms/Heading.vue';
import Paragraph from '@/components/atoms/Paragraph.vue';
import TitleHeader from '@/components/atoms/TitleHeader.vue';
import RegisterComponent from '@/components/organisms/RegisterComponent.vue';
</script>

<template>
  <main class="flex h-screen grid-cols-2 md:grid">
    <div
      class="hidden flex-1 flex-col items-center justify-center gap-4 text-center md:flex md:px-20"
    >
      <h1 class="font-libertinus text-5xl tracking-wide">The Tactile Breath</h1>
      <p class="font-manrope text-[12px] tracking-[0.2em] text-[#414846] uppercase">
        A journey into stillness begins here.
      </p>
    </div>
    <div class="flex flex-1 flex-col px-8 py-12">
      <div class="flex flex-1 flex-col justify-around lg:px-20">
        <TitleHeader class="text-4xl md:text-2xl" />
        <div>
          <h1 class="font-libertinus mb-8 text-5xl">Create Account</h1>
          <p class="font-manrope tracking-widest text-[#414846]">
            Join our curated community of wellness and international living.
          </p>
        </div>
        <RegisterComponent />
        <div class="flex flex-row justify-between">
          <Heading title="Already have an account?" headingClass="text-lg" />
          <Paragraph
            @click="$router.push('/sign-in')"
            title="Sign In "
            paraClass="font-bold cursor-pointer"
          />
        </div>
      </div>
    </div>
  </main>
</template>
