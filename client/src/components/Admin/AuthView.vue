<template>
    <div class="auth-inner">

        <div class="auth-wrapper">
            <h2 class="heading">
                Authorization
            </h2>

            <Section>
                <div class="input-wrapper">
                    <label class="input-label">
                        Email Address
                    </label>
                    <TextInput  v-model:data="email" 
                                :type="InputDataTypes.Email"
                                placeholder="teacher@example.com"
                                :required="true" />
                </div>

                <div class="input-wrapper">
                    <label class="input-label">
                        Password
                    </label>
                
                    <TextInput  v-model:data="password"
                                :type="InputDataTypes.Password"
                                placeholder="Qwerty1234 for example"
                                :required="true" />
                </div>
                
                <Button label="Login as Teacher"
                        class="internal-button"
                        @click="handleAuth" />

            </Section>
        </div>

    </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { loginTeacher } from '../../services/auth/login-service';
import { checkAuth } from '../../services/auth/check';
import Section from '../UI/Section.vue';
import TextInput from '../UI/TextInput.vue';
import { InputDataTypes } from '../../types/input-data-types';
import Button from '../UI/Button.vue';

export default defineComponent({
    name: "AuthView",
    components: {
        Section,
        TextInput,
        Button,
    },
    setup() {
        const router = useRouter();

        const email = ref('');
        const password = ref('');

        onMounted(async () => {
            const isValid = await checkAuth();
            if (isValid) {
                router.push({ name: "files" });
            }
        });

        const handleAuth = async () => {
            if (email.value === '' || password.value === '') {
                alert("Fill the required fields 🔍");
                return;
            }

            try {
                const res = await loginTeacher({
                    email: email.value,
                    password: password.value,
                });

                if (!res || !res.token) {
                    throw new Error("Token is missing in response");
                }

                localStorage.setItem("token", res.token);

                alert("Successfully login as a teacher ✅\n\nYou will be redirected to File Panel.");
                await router.push({ name: "files" });
            } catch (err) {
                console.error("Login error:", err);
                alert("Failed to login as a teacher ❌\n\nPlease try your credentials and try again.");
            }
        };

        return {
            InputDataTypes,
            handleAuth,
            email,
            password,
        };
    },
});
</script>

<style lang="scss" scoped>
@use "/src/assets/styles/colors" as *;
@use "/src/assets/styles/image";
@use "/src/assets/styles/fonts";
@use "/src/assets/styles/shadows";

.auth-inner {
    width: calc(100% - 30px - 30px);
    padding: 30px;
    height: calc(100vh - 72px);
    display: flex;
    justify-content: center;
    align-items: center;
}

.auth-wrapper {
    width: 90%;
    max-width: 500px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    
    .heading {
        @include fonts.noto-font(600);
        @include fonts.responsive-font(24, 20, 1440);
        text-align: center;
    }
}

.input-wrapper {
    width: 100%;
    max-width: 90vw;
    display: flex;
    flex-direction: column;
    gap: 4px;
    justify-content: space-between;
    align-items: flex-start;

    .input-label {
        @include fonts.noto-font(500);
        @include fonts.responsive-font(16, 16, 1440);
        text-align: start;
    }
}

.internal-button {
    margin-top: 10px;
}
</style>