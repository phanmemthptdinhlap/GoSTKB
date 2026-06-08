const ModalView ={
  template: `
  <div :class="theme.modal" v-if="showModal" >
    <div :class="theme.modal-container">
      <div :class="theme.modal_header">
        <div :class="theme.modal_title"> {{ modalTitle }} </div>
      </div>

      <div :class="theme.modal_body">
        <div :class="theme.modal_form">
          <div v-for="(label,index) in labels" :key="label.key" :class="theme.modal_item" >
            <label :class="theme.modal_item_label" > {{ label.label }}: </label>
            <input :class="theme.modal_item_input" type="text" v-model="data[label.key]">
          </div>
        </div>
        <div :class="theme.modal_button">
          <button :class="[theme.modal_button_save]" @click="saveData" :disabled="isLoading">
            {{ isLoading ? 'Đang lưu ... ' : 'Lưu' }}
          </button>
          <button :class="[theme.modal_button_cancel]" @click="closeModal" :disabled="isLoading">
            Hủy
          </button>
        </div>
      </div>
      <div>
        <div :class="theme.modal_message">{{ message }}</div>
      </div>
  </div>
</div>
    `,
  props : {
    labels: { type: Array, default: () => [] },
  },
  emits : ['saveData'],

  setup(props, context){
    const {ref} = Vue;
    const {emit, expose} = context;
    // ... (Giữ nguyên khai báo biến và hàm của bạn) ...
    const showModal = ref(false);
    const isEdit = ref(false);
    const isLoading = ref(false);
    const modalTitle = ref('Thêm mới');
    const data = ref({});
    const message = ref('');
      
    const closeModal = () => { showModal.value = false; isEdit.value = false; data.value = {}; };
    const openEditModal = (dataEdit,title) => { showModal.value = true; isEdit.value = true; modalTitle.value = title||'Chỉnh sửa'; data.value = {...dataEdit}; };
    const openAddModal = (title) => { showModal.value = true; isEdit.value = false; modalTitle.value = title||'Thêm mới'; data.value = {}; };
    const saveData = () => {
      isLoading.value = true;
      emit('saveData', data.value,(res) => {
        isLoading.value = false;
        if(res.success){ closeModal(); } else { message.value = res.message; }
        });
      };

    expose({ openEditModal, openAddModal, closeModal });
    
    return {
      theme: AppTheme, // QUAN TRỌNG: Phải return theme ở đây để template dùng được
      showModal,
      isEdit,
      isLoading,
      modalTitle,
      data,
      message,
      saveData,
      closeModal,
    };
  }
};
