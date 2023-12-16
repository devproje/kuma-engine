plugins {
    kotlin("jvm") version "1.9.21"
}

group = "net.projecttl"
version = "2.0.0"

val jda = property("jda_version")
val mongodb = property("mongodb_version")
val coroutine = property("coroutine_version")

allprojects {
    apply(plugin = "org.jetbrains.kotlin.jvm")

    kotlin {
        jvmToolchain(17)
    }

    repositories {
        mavenCentral()
    }
}


subprojects {
    repositories {
        mavenCentral()
    }

    dependencies {
        implementation(kotlin("stdlib"))
        implementation("net.dv8tion:JDA:${jda}")
        implementation("com.google.code.gson:gson:2.10.1")
        implementation("io.github.cdimascio:dotenv-kotlin:6.4.1")
        implementation("org.apache.logging.log4j:log4j-slf4j-impl:2.19.0")
        implementation("org.mongodb:mongodb-driver-kotlin-coroutine:${mongodb}")
        implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:${coroutine}")

        testImplementation(kotlin("test"))
    }
}
