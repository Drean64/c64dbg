
const SERVER = "http://localhost:6464";

async function _() {
	const res = await fetch(SERVER + "/state");
	res.formData()
		.then(fd => {
			for ([k, v] of fd) {
				console.log(k, v);
			}
		})
		.catch(err => console.log("Response.formData(): ", err));
	return new Promise((resolve) => resolve());
}

document.addEventListener('DOMContentLoaded', async () => {
	console.log("DOMContentLoaded");
	_();
	// const state = await fetch(SERVER + "/state");
	// const formData = await state.formData();
	// formData.forEach((value, key) => {
	// 	//
	// });
});