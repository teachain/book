<template>
    <Card>
        <Form :model="formItem">
            <FormItem label="Input" :label-width="80">
                <Input v-model="formItem.input" placeholder="Enter something..." />
            </FormItem>
            <FormItem label="Select" :label-width="80">
                <Select v-model="formItem.select">
                    <Option value="beijing">New York</Option>
                    <Option value="shanghai">London</Option>
                    <Option value="shenzhen">Sydney</Option>
                </Select>
            </FormItem>
            <FormItem :label-width="0">
                <Button type="primary" v-on:click="onSave">保存</Button>
            </FormItem>
        </Form>
    </Card>
</template>
<script>
export default {
    name: 'home',
    data() {
        return {
            formItem: {
                input: "",
                select: "",
            }
        }
    },
    methods: {
        onSave() {
            this.$Message.error('This is a info tip');
        }
    }
}
</script>