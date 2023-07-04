'use client';
import { download } from '@/lib/requests';
import styles from './table.module.css';
import { useRouter } from 'next/navigation';

type Files = {
  file: string;
  size: string;
};

export default function Table() {
  // const router = useRouter;
  // const data = download(router);

  return (
    <section className={styles.container}>
      <table className={styles.table}>
        {/* header row */}
        <tr className={styles.tableHeader}>
          <th>✔️</th>
          <th>File</th>
          <th>Size</th>
        </tr>
        {/* add row for each file  */}
        <tr className={styles.item}>
          <td>
            <input type='checkbox' />
          </td>
          <td>test.txt</td>
          <td>100kb</td>
        </tr>
      </table>
    </section>
  );
}

// 🡅 🡇
