import '../../styles/globals.css'
import 'tailwindcss/tailwind.css';
import type { AppProps } from 'next/app'
import AuthProvider from '../auth';
import axios from 'axios';

axios.defaults.baseURL = 'http://localhost:4000'

function MyApp({ Component, pageProps }: AppProps) {
  return (
 <AuthProvider>
  <Component {...pageProps} />
 </AuthProvider>
  )
}

export default MyApp
