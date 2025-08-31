import { ADMIN_EMAIL, ADMIN_ID, ADMIN_NAME, ADMIN_ROLE } from '@/models/auth';

const DashboardV2Page = () => {

  return (
    <div>

       <div>{localStorage.getItem(ADMIN_NAME)}</div>
       <div>{localStorage.getItem(ADMIN_EMAIL)}</div>
       <div>{localStorage.getItem(ADMIN_ID)}</div>
       <div>{localStorage.getItem(ADMIN_ROLE)}</div>
    </div>
  )
}

export default DashboardV2Page