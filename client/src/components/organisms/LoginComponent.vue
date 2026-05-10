<script setup lang="ts">
import InputGroup from '../molecules/InputGroup.vue';
import IconEyeOpen from '../../../public/icons/IconEyeOpen.vue';
import IconEyeClose from '../../../public/icons/IconEyeClose.vue';
import Button from '../atoms/Button.vue';
import Heading from '../atoms/Heading.vue';
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { signInWithEmailAndPassword } from 'firebase/auth';
import { auth } from '@/lib/firebase';
import { queryClient } from '@/query';

const router = useRouter();

const seePassword = ref(false);
const email = ref('');
const password = ref('');
const error = ref('');
const loading = ref(false);

const handleLogin = async () => {
  error.value = '';
  loading.value = true;

  try {
    // 1. Sign in with Firebase
    const userCredential = await signInWithEmailAndPassword(auth, email.value, password.value);
    const token = await userCredential.user.getIdToken();

    // 2. Verify with your Go server
    const response = await fetch('http://localhost:8080/users/me', {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      error.value = 'Failed to authenticate with server.';
      return;
    }

    const user = await response.json();

    // 3. Cache the user in TanStack Query
    queryClient.setQueryData(['auth'], userCredential.user);
    queryClient.setQueryData(['user'], user);

    // 4. Redirect to dashboard
    router.push('/');
  } catch (err: any) {
    if (err.code === 'auth/invalid-credential') {
      error.value = 'Invalid email or password.';
    } else if (err.code === 'auth/user-not-found') {
      error.value = 'No account found with this email.';
    } else if (err.code === 'auth/wrong-password') {
      error.value = 'Wrong password.';
    } else if (err.code === 'auth/too-many-requests') {
      error.value = 'Too many attempts. Please try again later.';
    } else {
      error.value = 'Something went wrong. Please try again.';
    }
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <form @submit.prevent="handleLogin" class="flex flex-col gap-10">
    <InputGroup
      type="email"
      id="email"
      label="Email"
      placeholder="ritual@sanctuary.com"
      @change="email = $event"
    />
    <InputGroup
      :type="seePassword ? 'text' : 'password'"
      placeholder="* * * * * * * *"
      @change="password = $event"
    >
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

    <p v-if="error" class="font-manrope text-sm text-red-500">{{ error }}</p>

    <Button :title="loading ? 'Signing in...' : 'Sign In'" type="submit" :disabled="loading" />
  </form>
</template>
