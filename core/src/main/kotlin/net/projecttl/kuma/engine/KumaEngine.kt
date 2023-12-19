package net.projecttl.kuma.engine

import kotlinx.coroutines.DelicateCoroutinesApi
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.coroutineScope
import kotlinx.coroutines.launch
import net.dv8tion.jda.api.JDABuilder
import net.dv8tion.jda.api.events.interaction.command.SlashCommandInteractionEvent
import net.dv8tion.jda.api.hooks.EventListener
import net.dv8tion.jda.api.requests.GatewayIntent
import net.dv8tion.jda.api.utils.cache.CacheFlag
import net.projecttl.kuma.engine.model.CommandExecutor
import net.projecttl.kuma.engine.handler.CommandHandler
import net.projecttl.kuma.engine.build.CommandDataBuilder
import net.projecttl.kuma.engine.build.NewEmbedBuilder
import org.slf4j.LoggerFactory
import java.io.File
import kotlin.random.Random

class KumaEngine(token: String, indents: List<GatewayIntent> = listOf(), flags: List<CacheFlag> = listOf()) {
    private val builder = JDABuilder.createDefault(token, indents)
        .enableCache(flags)
    private val commands = mutableListOf<CommandHandler>()
    private val logger = LoggerFactory.getLogger(KumaEngine::class.java)
    val command = CommandHandler()

    fun addHandler(vararg handler: EventListener) {
        handler.forEach {
            if (it is CommandHandler) {
                commands.add(it)
            }

            builder.addEventListeners(it)
        }

    }

    fun dropHandler(vararg handler: EventListener) {
        handler.forEach {
            if (it is CommandHandler) {
                commands.remove(it)
            }

            builder.removeEventListeners(it)
        }

    }

    companion object {
        val VERSION: String = File(this::class.java.getResource("/version.txt")!!.toURI()).readText()
    }

    @OptIn(DelicateCoroutinesApi::class)
    suspend fun build(info: Boolean = true) {
        addHandler(command)
        if (info) {
            addHandler(KumaInfo)
        }

        coroutineScope {
            launch {
                val jda = builder.build()
                GlobalScope.launch {
                    commands.forEach {
                        it.register(jda, logger)
                    }
                }
            }
        }
    }

    private object KumaInfo : CommandHandler() {
        private const val LOGO = "https://github.com/devproje/kuma-engine/raw/master/assets/kuma-engine-logo.png"
        init {
            addCommands(KumaInfoCommand)
        }

        private object KumaInfoCommand : CommandExecutor {
            override val data = CommandDataBuilder().apply {
                name = "kumainfo"
                description = "kuma engine system information"
            }

            override fun execute(event: SlashCommandInteractionEvent) {
                val embed = NewEmbedBuilder().apply {
                    title = ":dart: **KumaInfo**"
                    description = "KumaEngine system information"
                    thumbnail {
                        url = LOGO
                    }

                    field {
                        name = ":electric_plug: **ENGINE VERSION**"
                        value = "`${VERSION}`"
                        inline = true
                    }

                    field {
                        name = ":page_facing_up: **KOTLIN VERSION**"
                        value = "`${KotlinVersion.CURRENT}`"
                        inline = true
                    }

                    field {
                        name = ":ping_pong: **API LATENCY**"
                        value = "`${event.jda.gatewayPing}ms`"
                        inline = true
                    }

                    field {
                        name = ":desktop: **OS**"
                        value = "`${System.getProperty("os.name").lowercase()}/${System.getProperty("os.arch")}`"
                        inline = true
                    }

                    field {
                        name = ":satellite: **BOT SERVER**"
                        value = "`${event.jda.guilds.size}`"
                        inline = true
                    }

                    field {
                        name = ":file_folder: **SYSTEM PID**"
                        value = "`${ProcessHandle.current().pid()}`"
                        inline = true
                    }

                    footer {
                        text = event.user.name
                        iconUrl = event.user.avatarUrl ?: event.user.defaultAvatarUrl
                    }

                    color = Random.nextInt(0xFFFFFF + 1)
                }.build()

                event.replyEmbeds(embed).queue()
            }
        }
    }
}
