<template>
  <form @submit.prevent="" class="flex flex-col gap-10">
    <div class="flex flex-col gap-1">
      <label
        for="email"
        class="font-manrope text-[10px] font-bold tracking-[0.25em] text-[#47645C]/60 uppercase"
        >Email</label
      >
      <input
        type="email"
        id="email"
        placeholder="ritual@sanctuary.com"
        class="font-manrope h-11 border-b-[1px] border-[#C1C8C4] text-sm text-[#8e9792] outline-none placeholder:text-[#1C1C18]/20"
      />
    </div>
    <div class="flex flex-col gap-1">
      <label
        for="password"
        class="font-manrope text-[10px] font-bold tracking-[0.25em] text-[#47645C]/60 uppercase"
        >Password</label
      >
      <div class="flex flex-row items-center">
        <input
          type="password"
          id="password"
          placeholder="Ex. @S***123"
          class="font-manrope h-11 w-full border-b-[1px] border-[#C1C8C4] text-sm text-[#8e9792] outline-none placeholder:text-[#1C1C18]/20"
        />
      </div>
    </div>
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
