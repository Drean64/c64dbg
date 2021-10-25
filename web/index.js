
const SERVER = "http://localhost:6464";
let RAM, IO, CPU;

async function fetchState() {
	const res = await fetch(SERVER + "/state");
	const parts = await res.formData();
	for ([key, part] of parts) {
		if (part instanceof File) {
			const buffer = await part.arrayBuffer();
			if (key == 'RAM') {
				RAM = new Uint8Array(buffer)
			}
			else {
				IO = new Uint8Array(buffer)
			}
		}
		else if (key == 'CPU') {
			CPU = JSON.parse(part);
		}
	}
}

document.addEventListener('DOMContentLoaded', async () => {
	console.log("DOMContentLoaded");
	fetchState();
	// const state = await fetch(SERVER + "/state");
	// const formData = await state.formData();
	// formData.forEach((value, key) => {
	// 	//
	// });
});