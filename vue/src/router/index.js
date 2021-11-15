import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

import Layout from '@/layout'

export const constantRoutes = [{
  path: '/login',
  component: () => import('@/views/login'),
  hidden: true
},

{
  path: '/404',
  component: () => import('@/views/404'),
  hidden: true
},

{
  path: '/people-detail',
  component: Layout,
  redirect: '/people-detail',
  children: [{
    path: '/people-detail',
    name: 'PeopleDetail',
    component: () => import('@/views/peopleDetail'),
    meta: {
      title: 'People Detail',
      icon: 'sellingMe'
    }
  }],
  hidden: true
},

{
  path: '/',
  component: Layout,
  redirect: '/Dashboard',
  children: [{
    path: 'dashboard',
    name: 'Dashboard',
    component: () => import('@/views/dashboard'),
    meta: {
      title: 'Dashboard',
      icon: 'sellingMe'
    }
  }]
},
{
  path: '/checkin',
  component: Layout,
  redirect: '/checkin',
  children: [{
    path: '',
    name: 'Checkin',
    component: () => import('@/views/checkin'),
    meta: {
      title: 'Checkin',
      icon: 'sellingBuy'
    }
  }]
}
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */
export const asyncRoutes = [
  {
    path: '/people',
    component: Layout,
    redirect: '/people/list',
    name: 'People',
    alwaysShow: true,
    meta: {
      roles: ['admin', 'manager'],
      title: 'People',
      icon: 'realestate'
    },
    children: [
      {
        path: 'list',
        name: 'PeopleList',
        component: () => import('@/views/peopleList'),
        meta: {
          title: 'All People',
          icon: 'sellingAll'
        }
      },{
        path: 'add',
        name: 'addPeople',
        component: () => import('@/views/addPeople'),
        meta: {
          title: 'Add People',
          icon: 'sellingAll'
        }
      }
    ]
  },
  {
    path: '/agents',
    component: Layout,
    redirect: '/agent/list',
    name: 'People',
    alwaysShow: true,
    meta: {
      roles: ['admin', 'manager'],
      title: 'Agents',
      icon: 'addRealestate'
    },
    children: [
      {
        path: 'list',
        name: 'agentList',
        component: () => import('@/views/agentList'),
        meta: {
          title: 'All Agents',
          icon: 'sellingAll'
        }
      },{
        path: 'add',
        name: 'addPeople',
        component: () => import('@/views/addAgent'),
        meta: {
          title: 'Add Agent',
          icon: 'sellingAll'
        }
      }
    ]
  },
  {
    path: '/branches',
    component: Layout,
    redirect: '/branches/list',
    name: 'Branches',
    alwaysShow: true,
    meta: {
      roles: ['admin'],
      title: 'Branches',
      icon: 'sellingBuy'
    },
    children: [
      {
        path: 'list',
        name: 'branchList',
        component: () => import('@/views/branchList'),
        meta: {
          title: 'All Branches',
          icon: 'sellingAll'
        }
      },{
        path: 'add',
        name: 'addBranch',
        component: () => import('@/views/addBranch'),
        meta: {
          title: 'Add Branch',
          icon: 'sellingAll'
        }
      }
    ]
  },
  {
    path: '/records',
    component: Layout,
    redirect: '/records',
    children: [{
      path: '',
      name: 'Records',
      component: () => import('@/views/recordList'),
      meta: {
        roles: ['admin', 'manager'],
        title: 'Records',
        icon: 'selling'
      }
    }]
  },
  {
    path: '*',
    redirect: '/404',
    hidden: true
  }
]

const createRouter = () => new Router({
  mode: 'history',
  scrollBehavior: () => ({
    y: 0
  }),
  routes: constantRoutes
})

const router = createRouter()

export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
