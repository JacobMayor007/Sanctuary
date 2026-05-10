import { createRouter, createWebHistory } from 'vue-router';
import { queryClient } from '@/query';
import LoginView from '@/app/views/LoginView.vue';
import RegisterView from '@/app/views/RegisterView.vue';
import ForgotPassword from '@/app/views/ForgotPassword.vue';
import { auth } from '@/lib/firebase';

const routing = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      component: () => import('@/app/views/NotFoundView.vue'),
    },
    {
      path: '/',
      name: 'dashboard',
      component: () => import('@/app/views/DashboardView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/sign-in',
      name: 'sign-in',
      component: LoginView,
    },
    {
      path: '/sign-up',
      name: 'sign-up',
      component: RegisterView,
    },
    {
      path: '/forgot-password',
      name: 'forgot-password',
      component: ForgotPassword,
    },
  ],
});

routing.beforeEach(async (to) => {
  // Get cached auth user from TanStack Query
  let user = queryClient.getQueryData(['auth']);

  // If not cached, check Firebase directly
  if (!user) {
    user = await new Promise((resolve) => {
      const unsubscribe = auth.onAuthStateChanged((firebaseUser) => {
        unsubscribe();
        resolve(firebaseUser);
      });
    });

    // Cache it in TanStack Query
    if (user) {
      queryClient.setQueryData(['auth'], user);
    }
  }

  const isAuthenticated = !!user;

  if (isAuthenticated && (to.name === 'sign-in' || to.name === 'sign-up')) {
    return { name: 'dashboard' };
  }

  if (!isAuthenticated && to.meta.requiresAuth) {
    return { name: 'sign-in' };
  }
});

export default routing;
