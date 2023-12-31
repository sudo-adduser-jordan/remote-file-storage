'use client';
import Link from 'next/link';
import styles from './/download.module.css';
import { logout } from '@/lib/requests';
import Table from '@/app/components/Table';

export default function DownloadPage() {
  return (
    <main className={styles.main}>
      <div className={styles.title}>Remote File Storage</div>

      <div className={styles.container}>
        <Table />
        {/* <form
          className={styles.form}
          action='/pages/confirmation.html'
          method='POST'
          encType='multipart/form-data'
        >
          Download Files:
          <input className={styles.input} type='file' name='files[]' multiple />
        </form> */}
        <Link href={'/protected/confirm'}>
          <button className={styles.button}> 🡇 Download</button>
        </Link>
        <div className={styles.links}>
          <div className={styles.logout}>
            <Link href={'/'} onClick={logout}>
              Logout
            </Link>
          </div>
          <div className={styles.upload}>
            <Link href={'/protected/upload'}>Upload</Link>
          </div>
        </div>
      </div>
    </main>
  );
}

// 🡅 🡇
