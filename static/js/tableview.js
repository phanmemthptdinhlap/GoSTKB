const TableView ={
  template: `
  <div class="tv-main">
      <div class="tv-header">
        <div class="tv-title"> [[ title ]] </div>
        <div class="tv-subtitle"> [[ subtitle ]] </div>
      </div>
      <div class="tv-body">
        <table class="tv-table" border="1" style="width: 100%; border-collapse: collapse; text-align: left;">
          <thead class="tv-table-header">
            <tr>
              <th v-for="label in itemLabels" :key="label.key" style="padding: 8px;"> [[ label.title ]] </th>
              <th v-if="showButton" style="padding: 8px;">Thao tác</th>
            </tr>
          </thead>
          <tbody class="tv-table-body">
            <tr v-for="(item,index) in itemDatas" :key="index">
              <td v-for="label in itemLabels" :key="label.key" style="padding: 8px;"> [[ item[label.key] ]] </td>
              <td v-if="showButton" class="tv-table-cell-edit-button" style="padding: 8px;">
                <button class="tv-button" @click="openEditModal(index)">
                  <span class="tv-button-text">Sửa</span>
                </button>
                <button class="tv-button" @click="DeleteItem(index)">
                  <span class="tv-button-text">Xóa</span>
                </button>
              </td>
            </tr>
          </tbody>
          <tfoot class="tv-table-footer">
            <tr>
              <td colspan="100%" style="padding: 8px;">
                <div class="tv-pagination">
                  <span class="tv-pagination-text">Tổng số: [[ itemDatas.length ]] </span>
                  <button class="tv-pagination-button" @click="openAddModal">Thêm mới</button>
                </div>
              </td>
            </tr>
          </tfoot>
        </table>
      </div>
      <div class="tv-footer">
        <div class="tv-footer-text"> [[ footer ]] </div>
      </div>
    </div>
  <div class="modal" v-if="showModal" style="position: fixed; top: 20%; left: 30%; background: #eee; padding: 20px; border: 1px solid #ccc;">
      <div class="modal-header">
        <div class="modal-title"> [[ isEdit? 'Sửa thông tin' : 'Thêm mới' ]] [[ modalTitle ]] </div>
        <div class="modal-close" @click="closeModal" style="cursor: pointer; color: red;">
          <span class="modal-close-text">X</span>
        </div>
      </div>

      <div class="modal-body">
        <div class="modal-body-form">
          <div v-for="(label,index) in itemLabels" :key="label.key" class="modal-body-form-item" style="margin-bottom: 10px;">
            <label class="modal-body-form-item-label" style="display:inline-block; width: 80px;"> [[ label.title ]] </label>
            <input class="modal-body-form-item-input" type="text" v-model="modalItem[label.key]">
          </div>
        </div>
        <div class="modal-body-button">
          <button class="modal-body-button-text" @click="closeModal">Hủy</button>
          <button class="modal-body-button-text" @click="saveModal">Lưu</button>
        </div>
      </div>
    </div>
    `,
  props: {
    title: { type: String, default: 'Danh sách' },
    subtitle: { type: String, default: '' },
    footer: { type: String, default: '' },
    itemLabels: { type: Array, default: () => [] },
    itemDatas: { type: Array, default: () => [] },
    showButton: { type: Boolean, default: false }
  },
  data(){
    return {
      showModal: false,
      isEdit: false,
      modalItem: {},
      modalTitle: '',
    }
  },
  methods: {
    openEditModal(index){
      this.isEdit = true;
      this.modalItem = {...this.itemDatas[index]};
      this.modalTitle = this.title;
      this.showModal = true;
    },
    openAddModal(){
      this.isEdit = false;
      this.modalItem = {};
      this.modalTitle = this.title;
      this.showModal = true;
    },
    closeModal(){
      this.showModal = false;
    },
    saveModal(){
      this.$emit('updateItem', this.modalItem);
      this.closeModal();
    },
    DeleteItem(index){
      if(confirm('Bạn có chắc muốn xóa?'))
        this.$emit('deleteItem', index);
    }
  }
};
