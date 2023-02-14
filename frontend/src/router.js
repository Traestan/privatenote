import Vue from 'vue'
import Router from 'vue-router'
import HomeComponent from '@/views/Home';
//posts 
import EditComponent from '@/components/post/Edit';
import CreateComponent from '@/components/post/Create';
import PostComponent from '@/components/post/Post';
import NotesComponent from '@/components/post/Index';

//user
import UserLoginComponent from '@/components/user/Login';
import UserLogoutComponent from '@/components/user/Logout';
import UserRegisterComponent from '@/components/user/Register';
import UserForgotComponent from '@/components/user/Forgot';
import UserProfilePassComponent from '@/components/user/ProfilePass';

import { server } from "@/utils/helper";

Vue.use(Router)

let router =  new Router({
  mode: 'history',
  routes: [
    { path: '/', redirect: { name: 'home' } },
    { path: '/home', name: 'home', component: HomeComponent },
    /**/
    { path: '/create', name: 'Create', component: CreateComponent,meta: {requiresAuth: true,}  },
    { path: '/edit/:id', name: 'Edit', component: EditComponent,meta: {requiresAuth: true,}  },
    { path: '/post/:id', name: 'Post', component: PostComponent,meta: {requiresAuth: true,}  },
    { path: '/list/url', name: 'Notes', component: NotesComponent,meta: {requiresAuth: true,}  },
    

    /*
    { path: '/create/category', name: 'CreateCategory', component: CreateCategoryComponent },
    { path: '/edit/category/:id', name: 'EditCategory', component: EditCategoryComponent },
    { path: '/category/:id', name: 'Category', component: CategoryComponent },
    { path: '/categories', name: 'Categories', component: CategoriesComponent },
    
    //shorturl
    { path: '/create/url', name: 'CreateUrl', component: CreateShorturlComponent,meta: {requiresAuth: true,} },
    { path: '/edit/url/:shorturl', name: 'EditShorturl', component: EditShorturlComponent,meta: {requiresAuth: true,}  },
    { path: '/stats/:shorturl', name: 'StatUrl', component: ViewShorturlComponent,meta: {requiresAuth: true,}  },
    { path: '/list/url', name: 'Shorturls', component: ShorturlComponent,meta: {requiresAuth: true,}  },
    { path: '/del/:shorturl', name: 'DelUrl', component: ShorturlComponent,meta: {requiresAuth: true,}  },
    */
    //profile
    { path: '/user/changepass', name: 'ProfilePass', component: UserProfilePassComponent,meta: {requiresAuth: true,} },

    //user
    { path: '/login', name: 'UserLogin', component: UserLoginComponent },
    { path: '/logout', name: 'UserLogout', component: UserLogoutComponent },
    { path: '/register', name: 'UserRegister', component: UserRegisterComponent },
    { path: '/forgot', name: 'UserForgot', component: UserForgotComponent },
  ]
});

router.beforeEach((to, from, next) => {
  if(to.matched.some(record => record.meta.requiresAuth)) {
      if (localStorage.getItem('jwt') == null) {
          next({
              path: '/login',
              params: { nextUrl: to.fullPath }
          });
      }else {
          let user = localStorage.getItem('user');
          next();
      }
  } else if(to.matched.some(record => record.meta.guest)) {
      if(localStorage.getItem('jwt') == null){
        next();
      }else{
        next({ name: 'userboard'});
      }
  }else {
      next() 
  }
})

export default router 