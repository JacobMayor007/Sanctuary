<script setup lang="ts">
import LoginComponent from '@/components/organisms/LoginComponent.vue';
</script>

<template>
  <main class="relative flex h-screen flex-col overflow-hidden p-8">
    <!-- The title -->
    <h1 class="font-libertinus px-4 text-2xl tracking-[0.2em] text-[#47645C]/80 uppercase">
      The Sanctuary
    </h1>
    <h1
      class="font-libertinus rotate-12 pointer-events-none absolute right-28 bottom-16 -z-10 text-[192px] text-[#1C1C18]/10 italic"
    >
      S
    </h1>
    <div
      class="relative z-10 flex flex-1 flex-col gap-20 px-4 pt-16 md:justify-center md:px-32 md:pt-0 lg:px-[250px] xl:px-[400px]"
    >
      <div class="flex flex-col gap-10">
        <h1 class="font-libertinus text-5xl tracking-wide text-center">Welcome Home</h1>
        <p class="font-manrope text-center text-[14px] tracking-widest text-[#1C1C18]/50 uppercase">
          Enter your sanctuary to continue
        </p>
      </div>
      <LoginComponent />
      <h2 class="font-manrope text-center text-sm tracking-widest text-[#1C1C18]/50 uppercase">
        NEW TO THE Sanctuary?
        <span class="font-manrope font-bold text-[#1C1C18] uppercase">Create an account</span>
      </h2>
    </div>
  </main>
</template>
