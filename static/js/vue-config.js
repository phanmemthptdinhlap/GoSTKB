// Cấu hình mặc định cho Toast
export const toastOptions = {
    position: "top-right",
    timeout: 3000,
    closeOnClick: true,
    pauseOnFocusLoss: true,
    pauseOnHover: true,
    draggable: true,
    draggablePercent: 0.6,
    showCloseButtonOnHover: false,
    hideProgressBar: false,
    closeButton: "button",
    icon: true,
    rtl: false
};

// Hàm tạo ứng dụng Vue với Toast
export function createVueApp(options = {}) {
    const app = Vue.createApp(options);
    app.use(window.VueToastification.default, toastOptions);
    app.config.globalProperties.$toast = window.VueToastification.useToast();
    return app;
}