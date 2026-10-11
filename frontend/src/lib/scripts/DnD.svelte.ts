import { onMount } from 'svelte';
import { on } from '$lib/tools/on';
import { SelectFromDrop } from '$lib/wailsjs/go/api/Api';
import { OnFileDrop } from '$lib/wailsjs/runtime/runtime';

class DnD {
	dragOver = $state(false);

	init() {
		onMount(() => {
			const dragOverOff = on(window, 'dragover', (e) => e.preventDefault());
			const dropOff = on(window, 'drop', (e) => {
				e.preventDefault();
				requestAnimationFrame(() => (this.dragOver = false));
			});
			const dragLeaveOff = on(document.body, 'dragleave', (e) =>
				e.preventDefault(),
			);

			const dragEnterOff = on(window, 'dragenter', (e) => {
				e.preventDefault();
				if (this.dragOver) return;

				this.dragOver = true;

				requestAnimationFrame(() => {
					const dropTarget: HTMLElement | null =
						document.querySelector('.drop-target');
					if (!dropTarget) return;

					let timer: ReturnType<typeof setTimeout>;

					const deOff = on(dropTarget, 'dragenter', () => {
						clearTimeout(timer);
					});

					const dlOff = on(dropTarget, 'dragleave', () => {
						timer = setTimeout(() => {
							console.log('timeout');
							this.dragOver = false;
							dlOff();
							deOff();
						}, 500);
					});
				});
			});

			return () => {
				dragEnterOff();
				dropOff();
				dragLeaveOff();
				dragOverOff();
			};
		});

		OnFileDrop((...args) => {
			SelectFromDrop(...args);
			this.dragOver = false;
		}, true);
	}
}

export default new DnD();
