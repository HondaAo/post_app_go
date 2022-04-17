import type { NextPage } from 'next'
import Head from 'next/head'
import { useEffect, useState } from 'react'
import styles from '../../styles/Home.module.css'
import Layout from '../components/Layouts'
import { PostProps } from '../interface'
import axios from 'axios'
import { PostComponent } from '../components/Post'
import { useRouter } from 'next/router'

const Home: NextPage = () => {
  const [ posts, setPosts ] = useState<PostProps[]>([])
  const router = useRouter();
  const page = router.query
  useEffect(() => {
  axios.get(`/api/posts`, {
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
