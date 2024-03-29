import type { NextPage } from 'next'
import Head from 'next/head'
import ButtonCounter from '../components/ButtonCounter';
import ButtonAPI from '../components/ButtonAPI';
import ButtonAPIDB from '../components/ButtonAPIDB'

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
          playground
        </h1>
        <div>
          <ButtonCounter
          />
        </div>
        <div>
          <ButtonAPI
          />
        </div>
        <div>
          <ButtonAPIDB
            route={"randMySQLDB"}
          />
        </div>
        <div>
          <ButtonAPIDB
            route={"randPostgresDB"}
          />
        </div>
        <div>
          <ButtonAPIDB
            route={"randMongoDB"}
          />
        </div>
      </main>
    </div>
  )
}

export default Home
