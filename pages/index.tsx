import Image from 'next/image'
import { Inter } from 'next/font/google'
import Head from 'next/head'
import Link from 'next/link'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <main
      className={`flex h-screen flex-col  justify-between p-10 ${inter.className}`}
    >
      <ul className='flex relative ml-0'>
        <li className='font-bold mr-4'><strong>Nav-Api</strong></li>
        <li className='font-bold mr-4'><Link href='https://nav.programnotes.cn'>Nav</Link></li>
        <li className='font-bold'><Link href='https://github.com/serverless-coding/frontend-nav'>Github</Link></li>
      </ul>
      <div className="relative flex place-items-center items-center before:absolute before:h-[300px] before:w-[480px] before:-translate-x-1/2 before:rounded-full before:bg-gradient-radial before:from-white before:to-transparent before:blur-2xl before:content-[''] after:absolute after:-z-20 after:h-[180px] after:w-[240px] after:translate-x-1/3 after:bg-gradient-conic after:from-sky-200 after:via-blue-200 after:blur-2xl after:content-[''] before:dark:bg-gradient-to-br before:dark:from-transparent before:dark:to-blue-700/10 after:dark:from-sky-900 after:dark:via-[#0141ff]/40 before:lg:h-[360px]">
        <Image
          className="relative items-center dark:drop-shadow-[0_0_0.3rem_#ffffff70] dark:invert"
          src="/next.svg"
          alt="Next.js Logo"
          width={720}
          height={640}
          priority
        />
      </div>
      <footer className='relative flex justify-start'><p>&copy; serverless</p></footer>
    </main >
  )
}
