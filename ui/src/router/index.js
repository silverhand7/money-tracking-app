import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginView.vue'),
    },
    {
      path: '/',
      name: 'home',
      alias: ['/index.html'],
      component: () => import('@/views/CategoryView.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/categories',
      name: 'categories',
      component: () => import('@/views/CategoryView.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/categories/create',
      name: 'categories.create',
      component: () => import('@/views/CategoryForm.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/wallets',
      name: 'wallets',
      component: () => import('@/views/WalletView.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/wallet/create',
      name: 'wallets.create',
      component: () => import('@/views/WalletForm.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/wallets/:walletId/transactions',
      name: 'wallets.transactions',
      component: () => import('@/views/TransactionsView.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/transactions/create',
      name: 'transactions.create',
      component: () => import('@/views/TransactionForm.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/user',
      name: 'user',
      component: () => import('@/views/UserView.vue'),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: () => import('@/views/NotFoundView.vue'),
      meta: {
        requiresAuth: true
      }
    }
  ]
})

router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth) {
    const user = localStorage.getItem('user');
    if (user) {
      // call API to get User
      next();
    } else {
      next('/login');
    }
  } else {
    next();
  }
});

export default router
