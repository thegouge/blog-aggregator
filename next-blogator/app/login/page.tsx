import Header from "@/components/Header";
import Footer from "@/components/Footer"
import LogInForm from '@/components/LogInForm'

export default function Login() {

  return (
    <div className="flex min-h-screen flex-col items-center  justify-between p-24">
      <Header />
      <main>
        <h2>Log In!</h2>
        <LogInForm />
      </main>
      <Footer />
    </div>
  )
}
