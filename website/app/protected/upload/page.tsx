import Link from 'next/link';
import styles from './/upload.module.css';
import { logout } from '@/lib/requests';

export default function UploadPage() {
  return (
    <main className={styles.main}>
      <div className={styles.title}>Remote File Storage</div>

      <div className={styles.container}>
        <form
          className={styles.form}
          action='/pages/confirmation.html'
          method='POST'
          encType='multipart/form-data'
        >
          Upload Files:
          <input className={styles.input} type='file' name='files[]' multiple />
        </form>
        <Link href={'/protected/confirm'}>
          <button className={styles.button}> 🡅 Upload</button>
        </Link>
        <div className={styles.links}>
          <div className={styles.logout}>
            <Link href={'/'} onClick={logout}>
              Logout
            </Link>
          </div>
          <div className={styles.download}>
            <Link href={'/protected/download'}>Download</Link>
          </div>
        </div>
      </div>
    </main>
  );
}

// 🡅
