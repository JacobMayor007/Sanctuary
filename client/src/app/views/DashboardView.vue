<script setup lang="ts">
import Paragraph from '@/components/atoms/Paragraph.vue';
import TitleHeader from '@/components/atoms/TitleHeader.vue';
import { auth } from '@/lib';
import { useRouter } from 'vue-router';
import { queryClient } from '@/query';

const router = useRouter();

const handleSignout = async () => {
  await auth.signOut();

  queryClient.clear();

  router.push('/sign-in');
};
</script>

<template>
  <main>
    <TitleHeader />
    <Paragraph
      title="Sign out"
      @click="handleSignout"
      para-class="cursor-pointer active:scale-95"
    />
  </main>
</template>
