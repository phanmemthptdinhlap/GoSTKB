const DataTable = {
  props: {
    labels: { type: Array, required: true }, // Sửa: Đổi Object thành Array
    datas: { type: Array, required: true },  // Sửa: Đổi Object thành Array
  },
  template: `
    <div class="panel">
      <table class="table" style="width: 100%; border-collapse: collapse;">
        <thead>
          <tr>
            <!-- Sửa: Thay col.key thành col.id -->
            <template v-for="col in labels" :key="col.id">
              <th v-if="col.type == 'text'" style="padding: 10px; border: 1px solid #ccc; text-align: center;">
                {{ col.title }}
              </th>
              <template v-else-if="col.type == 'group'">
                <th v-for="cell in col.children" :key="col.id + '-' + cell.id" style="padding: 10px; border: 1px solid #ccc; text-align: center;">
                  {{ cell.title }}
                </th>
              </template>
            </template>
          </tr>
        </thead>
        <tbody>
          <!-- Thêm fallback cho key để tránh lỗi nếu row.id không tồn tại -->
          <tr v-for="row in localDatas" 
              :key="row.lop_id || row.id"
              style="border-bottom: 1px solid #eee; cursor: pointer;"
              @click="rowClick(row)">
            
            <template v-for="col in labels" :key="col.id">
              <td v-if="col.type == 'text'" style="padding: 8px; text-align: center; border: 1px solid #ccc;">
                {{ col.options[row[col.id]] || "-" }}
              </td>
              <template v-else-if="col.type == 'group'">
                <td v-for="cell in col.children" :key="col.id + '-' + cell.id"
                    style="padding: 8px; text-align: center; color: #d32f2f; border: 1px solid #ccc;">
                  <!-- Sửa: Đổi cell.options thành col.options -->
                  {{ col.options[row[col.id]?.[cell.id]] || '-' }}
                </td>
              </template>
            </template> 
            
          </tr>
        </tbody>
      </table>
    </div>
  `,
  setup(props) {
    const { ref, watch } = Vue; // Bổ sung watch từ Vue
    const localDatas = ref(props.datas);
    
    // Sửa: Sử dụng watch để đồng bộ dữ liệu hai chiều khi Datas bị thay đổi từ Component cha
    watch(() => props.datas, (newVal) => {
      localDatas.value = newVal;
    }, { immediate: true, deep: true });

    const rowClick = (row) => {
      alert(JSON.stringify(row));
    };

    return {
      localDatas,
      rowClick
    };
  } 
};
