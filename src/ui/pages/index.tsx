import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import Button from '../components/Button';
import styles from '../styles/Home.module.css'

const Home: NextPage = () => {
  return (
    <div className={styles.container}>
      <Head>
        <title>playground</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>
          Welcome
        </h1>
        <div>
          <Button 
          //randomNumber="0"
          />
        </div>
      </main>
    </div>
  )
}

export default Home
