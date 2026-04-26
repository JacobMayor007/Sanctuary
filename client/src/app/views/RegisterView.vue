<script setup lang="ts">
import Button from '@/components/atoms/Button.vue';
import Heading from '@/components/atoms/Heading.vue';
import Title from '@/components/atoms/title.vue';
import InputGroup from '@/components/molecules/InputGroup.vue';
</script>

<template>
  <main class="flex h-screen grid-cols-2 md:grid">
    <div class="hidden flex-1 flex-col items-center justify-center gap-4 md:flex">
      <h1 class="font-libertinus text-5xl tracking-wide">The Tactile Breath</h1>
      <p class="font-manrope text-[12px] tracking-[0.2em] text-[#414846] uppercase">
        A journey into stillness begins here.
      </p>
    </div>
    <div class="flex flex-1 flex-col px-8 py-12">
      <div class="flex flex-col gap-12">
        <Title />
        <div>
          <h1 class="font-libertinus mb-8 text-5xl">Create Account</h1>
          <p class="font-manrope tracking-widest text-[#414846]">
            Join our curated community of wellness and international living.
          </p>
        </div>
        <form action="" class="flex flex-col gap-5">
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
        <div>
          <Heading title="Already have an account?" headingClass="text-lg" />
        </div>
      </div>
    </div>
  </main>
</template>
