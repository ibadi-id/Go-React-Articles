import React from 'react';
const Welcome = React.lazy(() => import('./components/Welcome'));
const Preview = React.lazy(() => import('./components/Preview'));
const AllPosts = React.lazy(() => import('./components/AllPosts'));
const AddNew = React.lazy(() => import('./components/AddNew'));
const EditPost = React.lazy(() => import('./components/EditPost'));



const routes = [
  { path: '/', exact: true, name: 'Home', component: Welcome, exact: true },
  { path: '/preview', name: 'Preview', component: Preview, exact: true },
  { path: '/allposts', name: 'All Posts', component: AllPosts, exact: true },
  { path: '/addnew', name: 'Add New Post', component: AddNew, exact: true },
  { path: '/edit/:id', name: 'Edit Post', component: EditPost , exact: true },
];

export default routes;
