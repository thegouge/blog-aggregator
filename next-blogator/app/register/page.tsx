import Header from "@/components/Header";
import Footer from "@/components/Footer"
import RegisterForm from '@/components/RegisterForm'

export default function Register() {

  return (
    <div className="flex min-h-screen flex-col items-center  justify-between p-24">
      <Header />
      <main>
        <h2>Register!</h2>
        <RegisterForm />
      </main>
      <Footer />
    </div>
  )
}
