<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { createUserWithEmailAndPassword } from 'firebase/auth';
import Button from '../atoms/Button.vue';
import InputGroup from '../molecules/InputGroup.vue';
import { auth } from '@/lib/firebase';

const router = useRouter();

const name = ref('');
const email = ref('');
const password = ref('');
const error = ref('');
const loading = ref(false);

const handleRegister = async () => {
  error.value = '';
  loading.value = true;

  try {
    // 1. Register in Firebase on client side
    const userCredential = await createUserWithEmailAndPassword(auth, email.value, password.value);
    const token = await userCredential.user.getIdToken();

    // 2. Send token + name to your Go server to save in DB
    const response = await fetch('http://localhost:8080/users/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`, // 👈 send token
      },
      body: JSON.stringify({
        name: name.value,
      }),
    });

    if (!response.ok) {
      const message = await response.text();
      error.value = message;
      return;
    }

    const user = await response.json();
    console.log('Registered user:', user);
    router.push('/login');
  } catch (err: any) {
    // Firebase error messages
    if (err.code === 'auth/email-already-in-use') {
      error.value = 'Email already in use.';
    } else if (err.code === 'auth/weak-password') {
      error.value = 'Password should be at least 6 characters.';
    } else {
      error.value = 'Something went wrong. Please try again.';
    }
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <form @submit.prevent="handleRegister" class="flex flex-col gap-5 2xl:gap-5">
    <InputGroup
      class="mb-5"
      parentClass="gap-4"
      label="full name"
      type="text"
      placeholder="Eleanor Thorne"
      @change="name = $event"
    />
    <InputGroup
      class="mb-5"
      parentClass="gap-4"
      label="Email Address"
      type="email"
      placeholder="eleanorthorne@gmail.com"
      @change="email = $event"
    />
    <InputGroup
      class="mb-5"
      parentClass="gap-4"
      label="password"
      type="password"
      placeholder="***********"
      inputClass="tracking-[0.2em]"
      @change="password = $event"
    />

    <p v-if="error" class="font-manrope text-sm text-red-500">{{ error }}</p>

    <Button
      type="submit"
      :title="loading ? 'Creating account...' : 'Begin My Ritual'"
      buttonClass="py-5 rounded-xl bg-[#5F7D74] text-white text-[16px]"
      :disabled="loading"
    />
  </form>
</template>
