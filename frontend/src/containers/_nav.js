const _nav =  [
  {
    _tag: 'CSidebarNavDropdown',
    name: 'Posts',
    route: '/',
    icon: 'cil-puzzle',
    _children: [
      {
        _tag: 'CSidebarNavItem',
        name: 'All Posts',
        to: '/allposts',
      },
      {
        _tag: 'CSidebarNavItem',
        name: 'Add New',
        to: '/addnew',
      },
      {
        _tag: 'CSidebarNavItem',
        name: 'Preview',
        to: '/preview',
      },
    ],
  },
  {
    _tag: 'CSidebarNavDivider',
    className: 'm-2'
  }
]

export default _nav
