<script setup lang="ts">
import Heading from '@/components/atoms/Heading.vue';
import Paragraph from '@/components/atoms/Paragraph.vue';
import TitleHeader from '@/components/atoms/TitleHeader.vue';
import LoginComponent from '@/components/organisms/LoginComponent.vue';
</script>

<template>
  <main class="relative flex h-screen flex-col overflow-hidden p-8">
    <!-- The title -->
    <TitleHeader class="px-4" />

    <!-- This is the S Overlay -->
    <h1
      class="font-libertinus pointer-events-none absolute right-28 bottom-16 -z-10 rotate-12 text-[192px] text-[#1C1C18]/10 italic"
    >
      S
    </h1>
    <div
      class="relative z-10 flex flex-1 flex-col justify-around px-4 pt-16 md:justify-center md:gap-20 md:px-32 md:pt-0 lg:px-[250px] xl:px-[400px] 2xl:px-[500px]"
    >
      <div class="flex flex-col gap-10">
        <Heading title="Welcome Home" headingClass=" text-center text-5xl tracking-wide" />
        <Paragraph
          title="Enter your sanctuary to continue"
          paraClass="text-center text-[14px] tracking-widest text-[#1C1C18]/50 uppercase"
        />
      </div>
      <LoginComponent />
      <h2 class="font-manrope text-center text-sm tracking-widest text-[#1C1C18]/50 uppercase"></h2>
      <div class="flex flex-row justify-between">
        <Heading title="NEW TO THE Sanctuary?" headingClass="text-lg uppercase text-[#1C1C18]/50" />
        <Paragraph
          @click="$router.push('/sign-up')"
          title="Sign Up"
          paraClass="font-bold cursor-pointer uppercase"
        />
      </div>
    </div>
  </main>
</template>
