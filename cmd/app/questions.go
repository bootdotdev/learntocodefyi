package main

type QuestionType string

const (
	SingleSelectQuestionType QuestionType = "singleselect"
	MultiSelectQuestionType  QuestionType = "multiselect"
	OrderingQuestionType     QuestionType = "ordering"
	CountryQuestionType      QuestionType = "country"
)

type Question struct {
	QuestionType         QuestionType
	SingleSelectQuestion *SingleSelectQuestion
	MultiSelectQuestion  *MultiSelectQuestion
	OrderingQuestion     *OrderingQuestion
	CountryQuestion      *CountryQuestion
}

type SingleSelectQuestion struct {
	Question string
	Answers  []string
}

type MultiSelectQuestion struct {
	Question string
	Answers  []string
}

type OrderingQuestion struct {
	Question string
	Answers  []string
}

type CountryQuestion struct {
	Question string
}

func getQuestions() []Question {
	return []Question{
		{
			QuestionType: CountryQuestionType,
			CountryQuestion: &CountryQuestion{
				Question: "Where do you live?",
			},
		},
		{
			QuestionType: SingleSelectQuestionType,
			SingleSelectQuestion: &SingleSelectQuestion{
				Question: "How old are you?",
				Answers: []string{
					"Under 18",
					"18-24",
					"25-34",
					"35-44",
					"45-54",
					"Over 54",
				},
			},
		},
		{
			QuestionType: SingleSelectQuestionType,
			SingleSelectQuestion: &SingleSelectQuestion{
				Question: "How much formal education have you received? (Doesn't need to be programming related)",
				Answers: []string{
					"Primary/elementary school",
					"Secondary school (High school)",
					"Some college/university study without earning a degree",
					"Associate degree",
					"Bachelor's degree",
					"Master's degree",
					"Doctoral degree",
				},
			},
		},
		{
			QuestionType: SingleSelectQuestionType,
			SingleSelectQuestion: &SingleSelectQuestion{
				Question: "How many years of paid professional programming experience do you have, if any?",
				Answers: []string{
					"0",
					"<1",
					"1-2",
					"2-3",
					"3-4",
					"4-5",
					"5-6",
					"6-7",
					"7-8",
					"8+",
				},
			},
		},
		{
			QuestionType: MultiSelectQuestionType,
			MultiSelectQuestion: &MultiSelectQuestion{
				Question: "Have you completed any certified informal education?",
				Answers: []string{
					"Physical Bootcamp",
					"Online Bootcamp (With live instruction)",
					"Online Bootcamp (Fully self-paced, no live instruction)",
					"University certificate program",
					"Online course",
				},
			},
		},
		{
			QuestionType: SingleSelectQuestionType,
			SingleSelectQuestion: &SingleSelectQuestion{
				Question: "How long did you spend learning to code before you landed your first programming job? Only count months where you spent at least 10 hours per week learning. If you don't have a programming job, how long have you been studying in total?",
				Answers: []string{
					"<3 months",
					"3-6 months",
					"6-9 months",
					"9-12 months",
					"12-18 months",
					"18-24 months",
					"2-3 years",
					"3-4 years",
					"4-6 years",
					"6+ years",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "Which do you feel have been the most effective resources for learning a new programming language or framework?",
				Answers: []string{
					"Official documentation",
					"Books",
					"YouTube videos",
					"Coding live streams (Twitch, etc.)",
					"Podcasts",
					"Blog posts and articles",
					"Video courses",
					"Interactive courses",
					"Text-based courses",
					"Coding challenges",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "Which do you feel have been the most effective resources for learning a new tool or library?",
				Answers: []string{
					"Official documentation",
					"Books",
					"YouTube videos",
					"Coding live streams (Twitch, etc.)",
					"Podcasts",
					"Blog posts and articles",
					"Video courses",
					"Interactive courses",
					"Text-based courses",
					"Coding challenges",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "Where do you get most of your programming news?",
				Answers: []string{
					"Twitter/X",
					"Reddit",
					"Hacker News",
					"YouTube",
					"Podcasts",
					"Newsletters",
					"Personal Blogs",
					"Blogging platforms",
					"LinkedIn",
					"Facebook",
					"Instagram",
					"Snapchat",
					"TikTok",
					"Other",
				},
			},
		},
		{
			QuestionType: SingleSelectQuestionType,
			SingleSelectQuestion: &SingleSelectQuestion{
				Question: "How much do you (or did you) hope to earn in your first full time programming job?",
				Answers: []string{
					"<$20,000 / year",
					"$20,000 - $40,000 / year",
					"$40,000 - $60,000 / year",
					"$60,000 - $80,000 / year",
					"$80,000 - $100,000 / year",
					">$100,000 / year",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "Which types of developer roles are you most interested in?",
				Answers: []string{
					"Front-end web",
					"Back-end web",
					"Full-stack web",
					"Mobile",
					"Desktop or enterprise applications",
					"DevOps / SysAdmin / SRE / Cloud Infrastructure",
					"Data scientist or machine learning specialist",
					"Data engineer",
					"Blockchain",
					"Database Administrator",
					"Embedded applications or devices",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "Which programming languages are you most interested in working with professionally?",
				Answers: []string{
					"JavaScript",
					"HTML/CSS",
					"Python",
					"SQL",
					"TypeScript",
					"Java",
					"C#",
					"C++",
					"C",
					"PHP",
					"Go",
					"Rust",
					"Ruby",
					"Kotlin",
					"Dart",
					"Lua",
					"Swift",
					"Scala",
					"OCaml",
					"Zig",
					"other",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "Which programming languages are you most interested in learning personally?",
				Answers: []string{
					"JavaScript",
					"HTML/CSS",
					"Python",
					"SQL",
					"TypeScript",
					"Java",
					"C#",
					"C++",
					"C",
					"PHP",
					"Go",
					"Rust",
					"Ruby",
					"Kotlin",
					"Dart",
					"Lua",
					"Swift",
					"Scala",
					"OCaml",
					"Zig",
					"other",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "How do you balance your time in regard to passive and active learning? (Passive learning is watching videos, reading books, etc. Active learning is building projects, coding challenges, etc.)",
				Answers: []string{
					"Almost entirely passive",
					"Almost entirely active",
					"Mostly passive",
					"Mostly active",
					"50/50",
				},
			},
		},
		{
			QuestionType: SingleSelectQuestionType,
			SingleSelectQuestion: &SingleSelectQuestion{
				Question: "Which of the following do you use first for help when you're stuck?",
				Answers: []string{
					"Official documentation",
					"Web Search (e.g. Google)",
					"AI Chat (e.g ChatGPT)",
					"Forum (e.g. Stack Overflow, Reddit)",
					"Video Search (e.g. YouTube)",
					"Live community (e.g. Discord, Slack)",
					"A personal friend, family member, or colleague",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "Which of the following have been most effective for networking with professional developers?",
				Answers: []string{
					"Local Meetups",
					"Online Meetups",
					"Conferences",
					"Slack communities",
					"Discord communities",
					"Twitter/X",
					"LinkedIn",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "How much money have you spent on all educational resources out of your own pocket?",
				Answers: []string{
					"$0",
					"$1-$500",
					"$500-$1,999",
					"$2,000-$9,999",
					"$10,000-$49,999",
					"$50,000+",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "How much money has your employer spent on educational resources for you?",
				Answers: []string{
					"$0",
					"$1-$500",
					"$500-$1,999",
					"$2,000-$9,999",
					"$10,000-$49,999",
					"$50,000+",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "Which do you find to be most distracting when you're trying to focus on code?",
				Answers: []string{
					"Social Videos (YouTube, TikTok, etc.)",
					"Live Streams (Twitch, etc.)",
					"Social Media (Twitter, Facebook, etc.)",
					"News (Reddit, Hacker News, etc.)",
					"Live Chat (Discord, Slack, etc.)",
					"Home Life (Family, Roommates, etc.)",
					"Video Games",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "What do you find most difficult about learning to code?",
				Answers: []string{
					"Finding the time",
					"Losing interest or getting bored",
					"Getting help with specific problems when I'm stuck",
					"Knowing what to learn next",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "Which operating system do you code on?",
				Answers: []string{
					"Windows",
					"Windows (WSL)",
					"MacOS",
					"Linux",
					"ChromeOS",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "Which do you prefer to code in?",
				Answers: []string{
					"Full IDE (Visual Studio, IntelliJ, etc.)",
					"Lightweight IDE (VS Code, Atom, etc.)",
					"Terminal Editor (Vim, Emacs, etc.)",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "How much time do you spend passively learning to code per week?",
				Answers: []string{
					"Full IDE (Visual Studio, IntelliJ, etc.)",
					"Lightweight IDE (VS Code, Atom, etc.)",
					"Terminal Editor (Vim, Emacs, etc.)",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "Which more accurately describes you?",
				Answers: []string{
					"I can follow tutorials, but I struggle to build my own applications",
					"I understand the fundamentals of programming and computer science, but I struggle to build my own applications",
					"I've don't find building my own applications to be particularly difficult",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "What are you (or were you) most worried about?",
				Answers: []string{
					"I don't know which languages or frameworks I should choose",
					"I don't know if I should spend more time on theory or practice",
					"I'm worried AI will replace programmers",
					"I'm worried I'm not smart enough to be a programmer",
					"I'm worried I'm too old to be a programmer",
					"I'm worried I don't have enough time in my day to study",
					"I'm worried it will take too long to learn to code",
					"I'm worried that it's too expensive to learn to code",
					"I'm worried that my location will make it hard to get a job",
					"I'm worried that even though I'm spending time, I'm not making progress",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "What are you (or were you) most worried about when choosing a learning resource?",
				Answers: []string{
					"I don't know if this resource will be hands-on enough",
					"I don't know if this resource will be too hard or too easy",
					"I don't know if this resource will be too expensive",
					"I don't know if this resource will be too time consuming",
					"I don't know if this resource will be too boring",
					"I don't know if this resource will have innacurate or misleading information",
					"I don't know if this resource will teach the concepts clearly and thoroughly enough",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "Which benefits do you value most in a programming job?",
				Answers: []string{
					"High salary",
					"Health insurance",
					"Remote work",
					"Flexible hours",
					"Equity in the company",
					"Fun environment (ping pong, free food, etc.)",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "For your first job, which are you (or were you) most interested in?",
				Answers: []string{
					"Fully remote, same country",
					"Fully remote, different country",
					"Hybrid",
					"In office",
				},
			},
		},
		{
			QuestionType: SingleSelectQuestionType,
			SingleSelectQuestion: &SingleSelectQuestion{
				Question: "For your first job, do you plan on applying (or did you) to in-person jobs as well?",
				Answers: []string{
					"Yes",
					"No",
				},
			},
		},
		{
			QuestionType: SingleSelectQuestionType,
			SingleSelectQuestion: &SingleSelectQuestion{
				Question: "Which are you most worried about when it comes to qualifying for a programming job?",
				Answers: []string{
					"I'm worried I can't build production-ready applications",
					"I'm worried I don't understand computer science fundamentals well enough",
				},
			},
		},
		{
			QuestionType: SingleSelectQuestionType,
			SingleSelectQuestion: &SingleSelectQuestion{
				Question: "Are you actively looking for a programming job?",
				Answers: []string{
					"Yes, seriously",
					"Yes, but casually",
					"No",
				},
			},
		},
		{
			QuestionType: SingleSelectQuestionType,
			SingleSelectQuestion: &SingleSelectQuestion{
				Question: "How many programming jobs do you apply to per week when you're looking?",
				Answers: []string{
					"0",
					"1-5",
					"5-10",
					"10-20",
					"20-40",
					"40-80",
					"80-160",
					"160+",
				},
			},
		},
		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "Where do you look for programming jobs?",
				Answers: []string{
					"Big job boards (LinkedIn, Indeed, etc.)",
					"Small niche job boards",
					"Slack/Discord communities",
					"Personal network",
					"Physical Meetups",
					"Job fairs",
					"Conferences",
					"Cold outreach (dms, emails, etc.)",
					"Recruiters",
				},
			},
		},

		{
			QuestionType: OrderingQuestionType,
			OrderingQuestion: &OrderingQuestion{
				Question: "What are you (or were you) most worried about with applying to for your first programming job?",
				Answers: []string{
					"I don't have the practical skills",
					"I don't have the theoretical knowledge",
					"I don't have good portfolio projects",
					"I don't have a formal degree",
					"I don't have any professional experience",
					"I don't have a network",
					"I don't have a good personal brand or social media presence",
					"I don't know the right languages, frameworks, or tools",
				},
			},
		},
	}
}
