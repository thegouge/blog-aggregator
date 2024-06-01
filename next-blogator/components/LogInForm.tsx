'use client'

export default function LogInForm() {

	async function logInUser(formData: FormData) {
		const response = await fetch("http://localhost:8000/v1/login", {
			method: "POST",
			body: JSON.stringify({
				name: formData.get("userName"),
				password: formData.get("Password")
			}),
		})
		const responseData = await response.json()
		console.log({ responseData })
	}

	return (
		<form action={logInUser} className="flex-column">
			<label For="userName" className="block">
				User Name:
				<input type="text" name="userName" className="block" />
			</label>
			<label For="password" className="block">
				Password:<input type="password" name="password" className="block" />
			</label>
			<input type="submit" value="Submit!" />
		</ form>
	)
}

