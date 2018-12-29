<template>
    <Card>
        <Form :model="formItem">
            <FormItem label="名称" :label-width="80">
                <Select v-model="formItem.device_id">
                    <Option v-for="item in devices" :value="item.id" :key="item.id">{{ item.name }}</Option>
                </Select>
            </FormItem>
            <FormItem label="型号" :label-width="80">
                <Select v-model="formItem.driver_id">
                    <Option v-for="item in drivers" :value="item.id" :key="item.id">{{ item.name }}</Option>
                </Select>
            </FormItem>
             <FormItem label="" :label-width="80">
                <Input v-model="formItem.customer" placeholder="请输入..." />
            </FormItem>
             <FormItem label="" :label-width="80">
                <Input v-model="formItem.workplace" placeholder="请输入..." type="textarea" :rows="2"/>
            </FormItem>
             <FormItem label="" :label-width="80">
                <Input v-model="formItem.remarks" placeholder="请输入..." type="textarea" :rows="2"/>
            </FormItem>

            <FormItem label="" :label-width="80">
                 <DatePicker v-model="formItem.start_time" type="datetime" format="yyyy-MM-dd HH:mm" placeholder="选择开始日期和时间" style="width: 100%d"></DatePicker>
            </FormItem>

            <FormItem label="" :label-width="80">
                <DatePicker v-model="formItem.end_time"  type="datetime" format="yyyy-MM-dd HH:mm" placeholder="选择结束日期和时间" style="width: 100%d"></DatePicker>
            </FormItem>
            <FormItem :label-width="0">
                <Button type="primary" v-on:click="onSave">保存</Button>
            </FormItem>
        </Form>
    </Card>
</template>
<script>
export default {
    data() {
        return {
            formItem: {
                device_id: 0,
                driver_id: 0,
            },
            models: [],
        }
    },
    created() {
        this.$http.get("/model/list").then((data) => {
            this.models = data.data.data;
        });
    },
    methods: {
        onSave() {
            if (this.formItem.name == "") {
                this.$Message.error('挖掘机名称不能为空');
                return;
            }
            if (this.formItem.model_id == 0) {
                this.$Message.error('请选择挖掘机型号');
                return;
            }
            this.$http.post('/device/add', { name: this.formItem.name, model_id: this.formItem.model_id }).then((data) => {
                if (data.data.code == 0) {
                    this.$Message.success("挖掘机信息保存成功");
                    this.$router.push({ path: "/" });
                } else {
                    this.$Message.error("挖掘机信息保存失败");
                }
            });
        }
    }
}
</script>