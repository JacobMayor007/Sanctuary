<template>
  <div class="flex min-h-screen flex-col items-center justify-center gap-4">
    <h1 class="text-6xl font-bold text-[#47645C]">404</h1>
    <p class="font-manrope text-sm tracking-widest text-[#47645C] uppercase">Page Not Found</p>
    <button
      class="font-manrope cursor-pointer rounded-sm border-[1px] px-8 py-4.5 text-sm font-bold tracking-widest text-[#47645C] uppercase active:scale-95"
      @click="$router.push({ name: 'dashboard' })"
    >
      Go Home
    </button>
  </div>
</template>
