import '../../styles/globals.css'
import 'tailwindcss/tailwind.css';
import type { AppProps } from 'next/app'
import AuthProvider from '../auth';

export const url = 'http://localhost:4000'

function MyApp({ Component, pageProps }: AppProps) {
  return (
 <AuthProvider>
  <Component {...pageProps} />
 </AuthProvider>
  )
}

export default MyApp
