import { createRouter, createWebHistory } from 'vue-router';
import { queryClient } from '@/query';
import LoginView from '@/app/views/LoginView.vue';
import RegisterView from '@/app/views/RegisterView.vue';
import ForgotPassword from '@/app/views/ForgotPassword.vue';

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
  const user = queryClient.getQueryData(['auth']); // 🔁 swap with your actual query key

  const isAuthenticated = !!user;

  if (to.name === 'register') {
    return { name: 'sign-up' };
  }

  if (isAuthenticated && to.name === 'sign-in') {
    return { name: 'dashboard' };
  }

  if (!isAuthenticated && to.meta.requiresAuth) {
    return { name: 'sign-in' };
  }
});

export default routing;
