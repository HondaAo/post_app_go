import type { NextPage } from 'next'
import Head from 'next/head'
import { useEffect, useState } from 'react'
import styles from '../../styles/Home.module.css'
import Layout from '../components/Layouts'
import { Post } from '../interface'
import axios from 'axios'
import { url } from './_app'
import { PostComponent } from '../components/Post'
import { useRouter } from 'next/router'

const Home: NextPage = () => {
  const [ posts, setPosts ] = useState<Post[]>([])
  const router = useRouter();
  const page = router.query
  console.log(page)
  useEffect(() => {
  axios.get(`${url}/api/posts`, {
    withCredentials: true
  })
  .then(res => {
   setPosts(res.data)
  }).catch(err => console.log(err))
  },[])
  return (
    <Layout title='Home Page'>
     {posts ? posts!.map(post => {
       return (
        <PostComponent key={post.id} post={post} />
       )
     }) : (
       <div>
         Loading.....
       </div>
     )}
    </Layout>
  )
}

export default Home
